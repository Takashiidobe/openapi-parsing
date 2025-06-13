package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"openapi-parsing/pkg/openapi"
	"openapi-parsing/pkg/specfinder"
	"openapi-parsing/pkg/steps"
)

func resourceProviderToFolderName(resourceProvider string) string {
	rp := strings.TrimPrefix(resourceProvider, "Microsoft.")
	rp = strings.ReplaceAll(rp, "-", "_")
	return strings.ToLower(rp)
}

func determineSDKVersion(resourceProvider string) string {
	switch resourceProvider {
	case "Microsoft.DocumentDB":
		return "v3"
	case "Microsoft.ApiManagement":
		return "v3"
	case "Microsoft.Batch":
		return "v3"
	case "Microsoft.ContainerService":
		return "v6"
	case "Microsoft.Compute":
		return "v6"
	case "Microsoft.DataFactory":
		return "v10"
	case "Microsoft.AppContainers":
		return "v3"
	case "Microsoft.ContainerInstance":
		return "v2"
	case "Microsoft.Authorization":
		return "v2"
	case "Microsoft.DataProtection":
		return "v3"
	case "Microsoft.Cdn":
		return "v2"
	case "Microsoft.Communication":
		return "v2"
	case "Microsoft.AppConfiguration":
		return "v2"
	case "Microsoft.ApplicationInsights":
		return "v1"
	case "Microsoft.AzureStackHCI":
		return "v2"
	case "Microsoft.Avs":
		return "v2"
	case "Microsoft.DataBox":
		return "v2"
	case "Microsoft.BillingBenefits":
		return "v2"
	case "Microsoft.AppService":
		return "v4"
	case "Microsoft.Storage", "Microsoft.KeyVault", "Microsoft.Network",
		"Microsoft.Resources", "Microsoft.ManagedIdentity", "Microsoft.Insights",
		"Microsoft.OperationalInsights", "Microsoft.ContainerRegistry":
		return ""
	default:
		return ""
	}
}

func generateJSONAndGo(specFile string) error {
	parts := strings.Split(specFile, "/")
	var rp string
	for _, p := range parts {
		if strings.HasPrefix(p, "Microsoft.") {
			rp = p
			break
		}
	}
	if rp == "" {
		return fmt.Errorf("no Microsoft.* resource provider found in path")
	}
	folderName := resourceProviderToFolderName(rp)
	jsonDir := filepath.Join("json", folderName)
	goDir := filepath.Join("go", folderName)
	if err := os.MkdirAll(jsonDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(goDir, 0o755); err != nil {
		return err
	}

	fmt.Printf("Generating JSON for %s to %s\n", rp, jsonDir)
	jsonCmd := exec.Command("autorest",
		fmt.Sprintf("--input-file=%s", specFile),
		"--v3",
		"--use:@autorest/modelerfour",
		"--output-artifact=openapi-document",
		"--modelerfour.lenient-model-deduplication=true",
		"--clear-output-folder=true",
		fmt.Sprintf("--output-folder=%s", jsonDir),
	)
	if out, err := jsonCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("JSON generation failed: %s", string(out))
	}
	fmt.Printf("Successfully generated JSON for %s\n", rp)

	fmt.Printf("Generating Go code for %s to %s\n", rp, goDir)
	goCmd := exec.Command("autorest",
		fmt.Sprintf("--input-file=%s", specFile),
		"--v3",
		"--use:@autorest/modelerfour",
		"--modelerfour.lenient-model-deduplication=true",
		"--clear-output-folder=true",
		fmt.Sprintf("--output-folder=%s", goDir),
		"--go",
	)
	if out, err := goCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("Go generation failed: %s", string(out))
	}
	fmt.Printf("Successfully generated Go code for %s\n", rp)
	return nil
}

func dependencyExample(parser *openapi.Parser, target string) []openapi.Op {
	ops := parser.Ops()
	return openapi.FindDependencies(ops, target)
}

func main() {
	targets := []struct {
		path       string
		apiVersion string
	}{
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/listKeys", "2024-01-01"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default", "2024-08-15"},
	}

	for _, t := range targets {
		fmt.Printf("\n--- Processing target: %s with API version: %s ---\n", t.path, t.apiVersion)
		specs, err := specfinder.Find(t.apiVersion, t.path)
		if err != nil {
			log.Fatalf("Failed to find spec for target '%s' with API version '%s': %v", t.path, t.apiVersion, err)
		}
		if len(specs) == 0 {
			log.Fatalf("No spec files found for target: %s with API version: %s", t.path, t.apiVersion)
		}
		specFile := specs[0]
		fmt.Printf("Found spec file: %s\n", specFile)

		if err := generateJSONAndGo(specFile); err != nil {
			log.Fatalf("Failed to generate JSON and Go code for spec file '%s': %v", specFile, err)
		}
		fmt.Printf("Successfully generated JSON and Go code\n")

		parts := strings.Split(specFile, "/")
		var rp string
		for _, p := range parts {
			if strings.HasPrefix(p, "Microsoft.") {
				rp = p
				break
			}
		}
		folderName := resourceProviderToFolderName(rp)
		jsonFilePath := filepath.Join("./json", folderName, "openapi-document.json")
		fmt.Printf("Creating parser for: %s\n", jsonFilePath)

		parser, err := openapi.NewParser(jsonFilePath)
		if err != nil {
			log.Fatalf("openapi parse: %v", err)
		}
		deps := dependencyExample(parser, t.path)

		sdkVersion := determineSDKVersion(rp)
		v := sdkVersion
		if v == "" {
			v = "no version suffix"
		}
		fmt.Printf("Determined SDK version for %s: %s\n", rp, v)

		rootStep := steps.GenerateSteps(deps, sdkVersion)
		versionDir := folderName
		if sdkVersion != "" {
			versionDir = filepath.Join(folderName, sdkVersion)
		}
		specsDir := filepath.Join("./specs", versionDir)
		if err := os.MkdirAll(specsDir, 0o755); err != nil {
			log.Fatalf("Failed to create specs directory '%s': %v", specsDir, err)
		}
		outputFilename := filepath.Join(specsDir, "crawler.yaml")
		fmt.Printf("Writing crawler steps to: %s\n", outputFilename)
		if err := steps.WriteStepTreeAndSteps(outputFilename, rootStep); err != nil {
			log.Fatalf("Failed to write crawler file '%s': %v", outputFilename, err)
		}
		fmt.Printf("Successfully wrote crawler file: %s\n", outputFilename)
	}
}
