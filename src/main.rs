#[allow(unused)]
use std::collections::BTreeMap;

use oas3::{
    Spec,
    spec::{FromRef, ObjectOrReference, Parameter, Response, SchemaTypeSet},
};
use serde_json::value::Value;

#[derive(Debug, Clone)]
struct Parser {
    spec: Spec,
}

#[derive(Debug, Clone)]
struct Op {
    path: String,
    params: Vec<(String, SchemaTypeSet)>,
    response_type: String,
}

impl Parser {
    pub fn new(path: &str) -> Self {
        let json = std::fs::read_to_string(path).unwrap();

        Self {
            spec: oas3::from_json(json).unwrap(),
        }
    }

    pub fn paths(&self) -> Vec<Op> {
        let paths = self.spec.paths.clone().unwrap();

        let mut res = vec![];

        for (k, v) in paths {
            // select only gets
            if k.ends_with("get") {
                let path = self
                    .spec
                    .paths
                    .as_ref()
                    .unwrap()
                    .get(k.strip_suffix(".get").unwrap())
                    .unwrap()
                    .extensions
                    .get("ms-metadata")
                    .unwrap()
                    .get("path")
                    .unwrap()
                    .as_str()
                    .unwrap()
                    .to_string();

                // resolve all useful params
                let operation = v.get.unwrap().clone();
                let params = operation
                    .parameters(&self.spec)
                    .unwrap()
                    .into_iter()
                    .map(|p| {
                        (
                            p.name,
                            p.schema
                                .unwrap()
                                .resolve(&self.spec)
                                .unwrap()
                                .schema_type
                                .unwrap(),
                        )
                    })
                    .collect();

                let response = operation
                    .responses(&self.spec)
                    .get("200")
                    .unwrap()
                    .clone()
                    .content
                    .get("application/json")
                    .unwrap()
                    .schema(&self.spec)
                    .unwrap();
                let response_type = response
                    .extensions
                    .get("ms-metadata")
                    .unwrap()
                    .get("name")
                    .unwrap()
                    .as_str()
                    .unwrap()
                    .to_owned();

                let op = Op {
                    path,
                    params,
                    response_type,
                };

                res.push(op);
            }
        }

        res
    }

    pub fn operations(
        &self,
    ) -> (
        Vec<String>,
        Vec<Parameter>,
        Vec<Response>,
        Vec<BTreeMap<String, Value>>,
    ) {
        let mut operation_ids = vec![];
        let mut parameters = vec![];
        let mut responses = vec![];
        let mut extensions = vec![];
        for (k, method, operation) in self.spec.operations() {
            operation_ids.push(operation.operation_id.clone().unwrap());
            extensions.push(operation.extensions.clone());

            for parameter in &operation.parameters {
                parameters.push(self.resolve(parameter))
            }
            if method == "GET" {
                if let Some(r) = &operation.responses {
                    for (k, response) in r {
                        responses.push(self.resolve(response))
                    }
                }
                break;
            }
        }
        (operation_ids, parameters, responses, extensions)
    }

    fn resolve<T: Clone + FromRef>(&self, reference: &ObjectOrReference<T>) -> T {
        match reference {
            ObjectOrReference::Ref { .. } => reference.resolve(&self.spec).unwrap(),
            ObjectOrReference::Object(o) => o.clone(),
        }
    }
}

// it looks like we can just iterate through extensions
fn main() {
    let parser = Parser::new("./storage-accounts.json");
    let paths = parser.paths();
    dbg!(paths);

    // I just want the path and the api-version from the extensions
    // let (operation_ids, parameters, responses, extensions) = parser.operations();
    // for extension in extensions {
    //     for (k, v) in extension {
    //         let responses = &v["OperationsList"]["responses"]["200"]["body"]["value"];
    //         println!("{:#?}", responses);
    //     }
    // }
}
