use std::cmp::Ordering;
#[allow(unused)]
use std::collections::BTreeMap;

use oas3::{Spec, spec::SchemaTypeSet};

#[derive(Debug, Clone)]
pub struct Parser {
    spec: Spec,
}

#[derive(Debug, Clone, PartialEq)]
pub struct Op {
    path: String,
    params: Vec<(String, SchemaTypeSet)>,
    client: String,
    method: String,
    response_type: String,
}

impl Eq for Op {}

impl Ord for Op {
    fn cmp(&self, other: &Self) -> Ordering {
        match self.path.cmp(&other.path) {
            Ordering::Equal => self.response_type.cmp(&other.response_type),
            non_eq => non_eq,
        }
    }
}

impl PartialOrd for Op {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
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
                let operation_id = operation.operation_id.clone().expect("No operation id");
                let split: Vec<_> = operation_id.split('_').collect();
                let client = format!("{}Client", split[0]);
                let method = format!("New{}Pager", split[1]);
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
                    client,
                    method,
                    path,
                    params,
                    response_type,
                };

                res.push(op);
            }
        }
        res.sort();

        res
    }
}
