// it looks like we can just iterate through extensions
use openapi_parsing::go_parser::parse;
use openapi_parsing::openapi_parser::Parser;

fn main() {
    let parser = Parser::new("./storage-accounts.json");
    let paths = parser.paths();
    dbg!(paths);

    dbg!(parse());
}
