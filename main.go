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

// resourceProviderToFolderName converts Microsoft.Storage -> storage etc.
func resourceProviderToFolderName(rp string) string {
	rp = strings.TrimPrefix(rp, "Microsoft.")
	rp = strings.ReplaceAll(rp, "-", "_")
	return strings.ToLower(rp)
}

// determineSDKVersion maps resource providers to their SDK version suffix.
func determineSDKVersion(rp string) string {
	switch rp {
	case "Microsoft.DocumentDB", "Microsoft.ApiManagement", "Microsoft.Batch", "Microsoft.AppContainers":
		return "v3"
	case "Microsoft.ContainerService", "Microsoft.Compute":
		return "v6"
	case "Microsoft.DataFactory":
		return "v10"
	case "Microsoft.ContainerInstance", "Microsoft.Authorization":
		return "v2"
	case "Microsoft.DataProtection":
		return "v3"
	case "Microsoft.Cdn", "Microsoft.Communication", "Microsoft.AppConfiguration", "Microsoft.AzureStackHCI", "Microsoft.Avs", "Microsoft.DataBox", "Microsoft.BillingBenefits":
		return "v2"
	case "Microsoft.ApplicationInsights":
		return "v1"
	case "Microsoft.AppService":
		return "v4"
	default:
		// services without major version bumps
		return ""
	}
}

// generateJSONAndGo invokes autorest to produce Go code and OpenAPI JSON.
func generateJSONAndGo(specFile string) error {
	parts := strings.Split(specFile, string(os.PathSeparator))
	var rp string
	for _, p := range parts {
		if strings.HasPrefix(p, "Microsoft.") {
			rp = p
			break
		}
	}
	if rp == "" {
		return fmt.Errorf("no Microsoft.* resource provider found in %s", specFile)
	}

	folder := resourceProviderToFolderName(rp)
	jsonDir := filepath.Join("json", folder)
	goDir := filepath.Join("go", folder)
	if err := os.MkdirAll(jsonDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(goDir, 0755); err != nil {
		return err
	}

	jsonCmd := exec.Command("autorest",
		"--input-file="+specFile,
		"--v3",
		"--use:@autorest/modelerfour",
		"--output-artifact=openapi-document",
		"--modelerfour.lenient-model-deduplication=true",
		"--clear-output-folder=true",
		"--output-folder="+jsonDir,
	)
	if err := jsonCmd.Run(); err != nil {
		return fmt.Errorf("json generation failed: %w", err)
	}

	goCmd := exec.Command("autorest",
		"--input-file="+specFile,
		"--v3",
		"--use:@autorest/modelerfour",
		"--modelerfour.lenient-model-deduplication=true",
		"--clear-output-folder=true",
		"--output-folder="+goDir,
		"--go",
	)
	if err := goCmd.Run(); err != nil {
		return fmt.Errorf("go generation failed: %w", err)
	}
	return nil
}

func main() {
	targets := []struct {
		Path       string
		APIVersion string
	}{
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/listKeys", "2024-01-01"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default", "2024-08-15"},
	}

	for _, t := range targets {
		fmt.Printf("\n--- Processing target: %s with API version: %s ---\n", t.Path, t.APIVersion)

		specs, err := specfinder.Find(t.APIVersion, t.Path)
		if err != nil || len(specs) == 0 {
			log.Fatalf("failed to find spec for %s: %v", t.Path, err)
		}
		specFile := specs[0]
		fmt.Printf("Found spec file: %s\n", specFile)

		if err := generateJSONAndGo(specFile); err != nil {
			log.Fatalf("autorest generation: %v", err)
		}

		var rp string
		for _, p := range strings.Split(specFile, string(os.PathSeparator)) {
			if strings.HasPrefix(p, "Microsoft.") {
				rp = p
				break
			}
		}
		folder := resourceProviderToFolderName(rp)
		jsonFile := filepath.Join("json", folder, "openapi-document.json")

		parser, err := openapi.NewParser(jsonFile)
		if err != nil {
			log.Fatalf("openapi parse: %v", err)
		}
		deps := openapi.FindDependencies(parser.Ops(), t.Path)

		sdkVersion := determineSDKVersion(rp)
		root := steps.GenerateSteps(deps, sdkVersion)

		var versionDir string
		if sdkVersion != "" {
			versionDir = filepath.Join(folder, sdkVersion)
		} else {
			versionDir = folder
		}
		specsDir := filepath.Join("specs", versionDir)
		if err := os.MkdirAll(specsDir, 0755); err != nil {
			log.Fatalf("mkdir: %v", err)
		}
		outFile := filepath.Join(specsDir, "crawler.yaml")
		if err := steps.WriteStepTreeAndSteps(outFile, root); err != nil {
			log.Fatalf("write steps: %v", err)
		}
		fmt.Printf("wrote crawler file: %s\n", outFile)
	}
}
