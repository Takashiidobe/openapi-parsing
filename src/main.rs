// it looks like we can just iterate through extensions
use openapi_parsing::{
    find_rest_spec::spec_finder,
    openapi_parser::{Op, Parser, find_dependencies},
    step_generator::{generate_steps, write_step_tree_and_steps_to_file},
};
use std::process::Command;
use std::path::Path;
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

fn main() {
    // let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts";
    let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/listKeys";

    let parser = Parser::new("./json/storage/openapi-document.json");
    let ex = dependency_example(parser, target);
    let package_version = "v3";

    let root_step = generate_steps(&ex, &package_version);

    write_step_tree_and_steps_to_file(root_step, "crawler_output.yaml").unwrap();

    let api_version = "2024-08-15";
    let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default";
    match spec_finder(api_version, target) {
        Ok(specs) => {
            if let Some(spec_file) = specs.first() {
                println!("Found spec file: {}", spec_file);
                
                // Generate JSON and Go code from the found spec file
                match generate_json_and_go(spec_file) {
                    Ok(()) => println!("Successfully generated JSON and Go code"),
                    Err(e) => eprintln!("Error generating code: {}", e),
                }
            } else {
                println!("No spec files found for the given criteria");
            }
        }
        Err(e) => eprintln!("Error finding spec: {}", e),
    }
}
