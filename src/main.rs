// it looks like we can just iterate through extensions
use openapi_parsing::{
    find_rest_spec::spec_finder,
    openapi_parser::{Op, Parser, find_dependencies},
    step_generator::generate_steps,
};

use openapi_parsing::go_parser::parse;

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

    dbg!(generate_steps(&ex));

    // let api_version = "2024-08-15";
    // let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default";
    // spec_finder(api_version, target).unwrap();
    // now that we have dependencies, I want to generate a StepTree for each op.
    // So Vec<Op> -> StepTree + Vec<Step>
}
