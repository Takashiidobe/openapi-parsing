#[allow(unused)]
use std::collections::BTreeMap;
use std::{cmp::Ordering, collections::HashSet};

use oas3::{Spec, spec::SchemaTypeSet};

#[derive(Debug, Clone)]
pub struct Parser {
    pub spec: Spec,
}

#[derive(Debug, Clone, PartialEq)]
pub struct Op {
    pub path: String,
    pub params: Vec<(String, SchemaTypeSet)>,
    pub client: String,
    pub method: String,
    pub response_type: String,
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

                // find the resource provider
                let split_path: Vec<_> = path.split("/").collect();

                let resource_provider = split_path
                    .iter()
                    .find(|s| s.starts_with("Microsoft."))
                    .expect("No Microsoft.* entry found")
                    .strip_prefix("Microsoft.")
                    .expect("Entry did not start with Microsoft.");

                let mut client = format!("{}Client", split[0]);
                client = client.strip_prefix(resource_provider).unwrap_or(&client).to_string();

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

fn prefixes(path: &str) -> Vec<String> {
    // trim leading slash so we don't get an empty first segment
    let trimmed = path.trim_start_matches('/');
    let mut acc = String::new();
    let mut out = Vec::new();
    for segment in trimmed.split('/') {
        acc.push('/');
        acc.push_str(segment);
        out.push(acc.clone());
    }
    out
}

pub fn find_dependencies(ops: &[Op], target_path: &str) -> Vec<Op> {
    // build a fast lookup of all valid prefixes
    let prefix_set: HashSet<_> = prefixes(target_path).into_iter().collect();

    // filter and then sort by path length so you get the roots first
    let mut deps: Vec<&Op> = ops
        .iter()
        .filter(|op| prefix_set.contains(&op.path))
        .collect();

    deps.sort_by_key(|op| op.path.len());
    deps.into_iter().cloned().collect()
}
