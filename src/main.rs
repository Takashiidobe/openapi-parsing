// it looks like we can just iterate through extensions
use openapi_parsing::{
    openapi_parser::{Op, Parser, find_dependencies},
    step_generator::generate_steps,
};

pub fn dependency_example(parser: Parser) -> Vec<Op> {
    let paths = parser.paths();

    let target = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default";

    find_dependencies(&paths, target)
}

fn main() {
    let parser = Parser::new("./json/cosmos_db/openapi-document.json");
    // dbg!(parse("./go/cosmos_db/").unwrap());
    let ex = dependency_example(parser);

    dbg!(generate_steps(&ex));
    // now that we have dependencies, I want to generate a StepTree for each op.
    // So Vec<Op> -> StepTree + Vec<Step>
}
