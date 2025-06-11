// it looks like we can just iterate through extensions
use openapi_parsing::{
    find_rest_spec::spec_finder,
    openapi_parser::{Op, Parser, find_dependencies},
    step_generator::{generate_steps, write_step_tree_and_steps_to_file},
};
use std::process::Command;
use std::fs;

pub fn dependency_example(parser: Parser, target: &str) -> Vec<Op> {
    let paths = parser.paths();

    find_dependencies(&paths, target)
}

fn resource_provider_to_folder_name(resource_provider: &str) -> String {
    // Convert Microsoft.DocumentDB -> cosmos_db, Microsoft.Storage -> storage, etc.
    resource_provider
        .strip_prefix("Microsoft.")
        .unwrap_or(resource_provider)
        .replace("-", "_")
        .to_lowercase()
}

/// Determine the SDK version based on the resource provider
/// This maps Azure service resource providers to their corresponding Go SDK versions
fn determine_sdk_version(resource_provider: &str) -> Option<&'static str> {
    match resource_provider {
        // Services with major version bumps
        "Microsoft.DocumentDB" => Some("v3"),  // CosmosDB uses v3
        "Microsoft.ApiManagement" => Some("v3"),  // API Management uses v3
        "Microsoft.Batch" => Some("v3"),  // Batch uses v3
        "Microsoft.ContainerService" => Some("v6"),  // Container Service uses v6
        "Microsoft.Compute" => Some("v6"),  // Compute uses v6
        "Microsoft.DataFactory" => Some("v10"),  // Data Factory uses v10
        "Microsoft.AppContainers" => Some("v3"),  // Container Apps uses v3
        "Microsoft.ContainerInstance" => Some("v2"),  // Container Instances uses v2
        "Microsoft.Authorization" => Some("v2"),  // Authorization uses v2
        "Microsoft.DataProtection" => Some("v3"),  // Data Protection uses v3
        "Microsoft.Cdn" => Some("v2"),  // CDN uses v2
        "Microsoft.Communication" => Some("v2"),  // Communication uses v2
        "Microsoft.AppConfiguration" => Some("v2"),  // App Configuration uses v2
        "Microsoft.ApplicationInsights" => Some("v1"),  // Application Insights still v1
        "Microsoft.AzureStackHCI" => Some("v2"),  // Azure Stack HCI uses v2
        "Microsoft.Avs" => Some("v2"),  // Azure VMware Solution uses v2
        "Microsoft.DataBox" => Some("v2"),  // Data Box uses v2
        "Microsoft.BillingBenefits" => Some("v2"),  // Billing Benefits uses v2
        "Microsoft.AppService" => Some("v4"),  // App Service uses v4
        
        // Services without major version bumps (use no version suffix)
        "Microsoft.Storage" => None,  // Storage doesn't have version suffix
        "Microsoft.KeyVault" => None,  // Key Vault doesn't have version suffix
        "Microsoft.Network" => None,  // Network typically doesn't have version suffix
        "Microsoft.Resources" => None,  // Resources doesn't have version suffix
        "Microsoft.ManagedIdentity" => None,  // Managed Identity doesn't have version suffix
        "Microsoft.Insights" => None,  // Insights doesn't have version suffix
        "Microsoft.OperationalInsights" => None,  // Operational Insights doesn't have version suffix
        "Microsoft.ContainerRegistry" => None,  // Container Registry doesn't have version suffix
        
        // Default case for unknown services - assume no version suffix
        _ => None,
    }
}

fn generate_json_and_go(spec_file: &str) -> Result<(), Box<dyn std::error::Error>> {
    // Extract resource provider from spec file path
    // e.g., "../azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json"
    let path_parts: Vec<&str> = spec_file.split('/').collect();
    let resource_provider = path_parts
        .iter()
        .find(|part| part.starts_with("Microsoft."))
        .ok_or("No Microsoft.* resource provider found in path")?;
    
    let folder_name = resource_provider_to_folder_name(resource_provider);
    
    // Create output directories
    let json_dir = format!("json/{}", folder_name);
    let go_dir = format!("go/{}", folder_name);
    
    fs::create_dir_all(&json_dir)?;
    fs::create_dir_all(&go_dir)?;
    
    println!("Generating JSON for {} to {}", resource_provider, json_dir);
    
    // Generate JSON (OpenAPI v3)
    let json_output = Command::new("autorest")
        .arg(format!("--input-file={}", spec_file))
        .arg("--v3")
        .arg("--use:@autorest/modelerfour")
        .arg("--output-artifact=openapi-document")
        .arg("--modelerfour.lenient-model-deduplication=true")
        .arg("--clear-output-folder=true")
        .arg(format!("--output-folder={}", json_dir))
        .output()?;
    
    if !json_output.status.success() {
        eprintln!("JSON generation failed: {}", String::from_utf8_lossy(&json_output.stderr));
        return Err("JSON generation failed".into());
    }
    
    println!("Successfully generated JSON for {}", resource_provider);
    
    println!("Generating Go code for {} to {}", resource_provider, go_dir);

    // Generate Go code
    let go_output = Command::new("autorest")
        .arg(format!("--input-file={}", spec_file))
        .arg("--v3")
        .arg("--use:@autorest/modelerfour")
        .arg("--modelerfour.lenient-model-deduplication=true")
        .arg("--clear-output-folder=true")
        .arg(format!("--output-folder={}", go_dir))
        .arg("--go")
        .output()?;
    
    if !go_output.status.success() {
        eprintln!("Go generation failed: {}", String::from_utf8_lossy(&go_output.stderr));
        return Err("Go generation failed".into());
    }
    
    println!("Successfully generated Go code for {}", resource_provider);
    
    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let targets = vec![
        ("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/listKeys", "2024-01-01"),
        ("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default", "2024-08-15"),
    ];

    for (target, api_version) in targets {
        println!("\n--- Processing target: {} with API version: {} ---", target, api_version);
        
        // First, find the spec file for this target
        let specs = spec_finder(api_version, target)
            .map_err(|e| format!("Failed to find spec for target '{}' with API version '{}': {}", target, api_version, e))?;
        
        let spec_file = specs.first()
            .ok_or_else(|| format!("No spec files found for target: {} with API version: {}", target, api_version))?;
        
        println!("Found spec file: {}", spec_file);
        
        // Generate JSON and Go code from the found spec file
        generate_json_and_go(spec_file)
            .map_err(|e| format!("Failed to generate JSON and Go code for spec file '{}': {}", spec_file, e))?;
        println!("Successfully generated JSON and Go code");
        
        // Extract resource provider to determine the JSON file path
        let path_parts: Vec<&str> = spec_file.split('/').collect();
        let resource_provider = path_parts.iter()
            .find(|part| part.starts_with("Microsoft."))
            .ok_or_else(|| format!("Could not extract resource provider from spec file path: {}", spec_file))?;
        
        let folder_name = resource_provider_to_folder_name(resource_provider);
        let json_file_path = format!("./json/{}/openapi-document.json", folder_name);
        
        println!("Creating parser for: {}", json_file_path);
        
        // Create parser from the generated JSON file
        let parser = Parser::new(&json_file_path);
        let ex = dependency_example(parser, target);
        
        // Determine SDK version based on the resource provider
        let sdk_version = determine_sdk_version(resource_provider);
        println!("Determined SDK version for {}: {}", resource_provider, 
            sdk_version.map_or("no version suffix".to_string(), |v| v.to_string()));
        
        let root_step = generate_steps(&ex, sdk_version.unwrap_or(""));
        
        // Create unique output filename based on resource provider and include version in directory structure
        let version_dir = match sdk_version {
            Some(version) => format!("{}/{}", folder_name, version),
            None => folder_name.clone(),
        };
        
        // Ensure the specs directory structure exists
        let specs_dir = format!("./specs/{}", version_dir);
        fs::create_dir_all(&specs_dir)
            .map_err(|e| format!("Failed to create specs directory '{}': {}", specs_dir, e))?;
        
        let output_filename = format!("{}/crawler.yaml", specs_dir);
        println!("Writing crawler steps to: {}", output_filename);
        
        write_step_tree_and_steps_to_file(root_step, &output_filename)
            .map_err(|e| format!("Failed to write crawler file '{}': {}", output_filename, e))?;
        println!("Successfully wrote crawler file: {}", output_filename);
    }
    
    Ok(())
}
