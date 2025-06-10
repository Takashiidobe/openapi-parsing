// it looks like we can just iterate through extensions
use openapi_parsing::{
    find_rest_spec::spec_finder,
    openapi_parser::{Op, Parser, find_dependencies},
    step_generator::{generate_steps, write_step_tree_and_steps_to_file},
};

pub fn dependency_example(parser: Parser) -> Vec<Op> {
    let paths = parser.paths();

    let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/listKeys";
    // let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts";

    find_dependencies(&paths, target)
}

fn main() {
    let parser = Parser::new("./json/storage/openapi-document.json");
    // dbg!(parse("./go/storage/").unwrap());
    let ex = dependency_example(parser);

    let root_step = generate_steps(&ex, "v3");

    // Write step tree and serialized steps to file
    match write_step_tree_and_steps_to_file(root_step, "crawler_output.yaml") {
        Ok(()) => println!("Successfully wrote step tree and steps to crawler_output.yaml"),
        Err(e) => eprintln!("Error writing to file: {}", e),
    }

    let api_version = "2024-08-15";
    let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default";
    dbg!(spec_finder(api_version, target).unwrap());
}
