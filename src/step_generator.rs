// Function to take Vec<Op> -> Step

use crate::openapi_parser::Op;
use serde::{Deserialize, Serialize};
use std::fs;

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash, Serialize, Deserialize)]
pub struct ClientMethod {
    package: String,
    client: String,
    method: String,
    args: Vec<String>,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash, Serialize, Deserialize)]
pub struct ChannelStep {
    pub id: String,
    client_method: ClientMethod,
    pub children: Vec<Step>,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash, Serialize, Deserialize)]
pub struct ResourceStep {
    pub id: String,
    resource: String,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash, Serialize, Deserialize)]
pub struct PayloadStep {
    pub id: String,
    exclude_tags: Vec<String>,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash, Serialize, Deserialize)]
pub struct StepTree {
    root_step: Step,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash, Serialize, Deserialize)]
pub enum Step {
    Channel(ChannelStep),
    Resource(ResourceStep),
    Payload(PayloadStep),
}

// Helper struct for YAML output format
#[derive(Serialize)]
struct StepNode {
    id: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    children: Option<Vec<StepNode>>,
}

#[derive(Serialize)]
struct Executor {
    default: serde_yaml::Value,
}

#[derive(Serialize)]
struct YamlOutput {
    kind: String,
    id: String,
    executor: Executor,
    #[serde(rename = "rootStep")]
    root_step: StepNode,
}

impl From<&Step> for StepNode {
    fn from(step: &Step) -> Self {
        match step {
            Step::Channel(channel) => {
                let children = if !channel.children.is_empty() {
                    Some(channel.children.iter().map(StepNode::from).collect())
                } else {
                    None
                };
                StepNode {
                    id: channel.id.clone(),
                    children,
                }
            }
            Step::Resource(resource) => StepNode {
                id: resource.id.clone(),
                children: None,
            },
            Step::Payload(payload) => StepNode {
                id: payload.id.clone(),
                children: None,
            },
        }
    }
}

// Structs for serializing individual steps
#[derive(Serialize)]
struct SerializedChannelStep {
    kind: String,
    id: String,
    #[serde(rename = "clientMethod")]
    client_method: SerializedClientMethod,
}

#[derive(Serialize)]
struct SerializedClientMethod {
    package: String,
    client: String,
    method: String,
}

#[derive(Serialize)]
struct SerializedResourceStep {
    kind: String,
    id: String,
    resource: String,
}

#[derive(Serialize)]
struct SerializedPayloadStep {
    kind: String,
    id: String,
}

// Helper function to convert camelCase to snake_case with azure_ prefix
fn to_azure_resource_name(name: &str) -> String {
    // Remove "ListResult" suffix if present to get singular resource name
    let clean_name = if name.ends_with("ListResult") {
        name.strip_suffix("ListResult").unwrap_or(name)
    } else {
        name
    };
    
    let mut result = String::from("azure_");
    let mut chars = clean_name.chars().peekable();
    
    while let Some(ch) = chars.next() {
        if ch.is_uppercase() && !result.ends_with("azure_") {
            result.push('_');
        }
        result.push(ch.to_lowercase().next().unwrap());
    }
    
    result
}

fn serialize_step_recursive(step: &Step) -> Vec<String> {
    let mut results = Vec::new();
    let schema_comment = "# yaml-language-server: $schema=../../../../../crawler-generator/crawler.schema.json\n";
    
    match step {
        Step::Channel(channel) => {
            let serialized = SerializedChannelStep {
                kind: "ChannelStep".to_string(),
                id: channel.id.clone(),
                client_method: SerializedClientMethod {
                    package: channel.client_method.package.clone(),
                    client: channel.client_method.client.clone(),
                    method: channel.client_method.method.clone(),
                },
            };
            
            if let Ok(yaml) = serde_yaml::to_string(&serialized) {
                results.push(format!("{}{}", schema_comment, yaml));
            }
            
            // Recursively serialize children
            for child in &channel.children {
                results.extend(serialize_step_recursive(child));
            }
        }
        Step::Resource(resource) => {
            let serialized = SerializedResourceStep {
                kind: "ResourceStep".to_string(),
                id: resource.id.clone(),
                resource: to_azure_resource_name(&resource.resource),
            };
            
            if let Ok(yaml) = serde_yaml::to_string(&serialized) {
                results.push(format!("{}{}", schema_comment, yaml));
            }
        }
        Step::Payload(payload) => {
            let serialized = SerializedPayloadStep {
                kind: "PayloadStep".to_string(),
                id: payload.id.clone(),
            };
            
            if let Ok(yaml) = serde_yaml::to_string(&serialized) {
                results.push(format!("{}{}", schema_comment, yaml));
            }
        }
    }
    
    results
}

pub fn serialize_steps(root_step: Step) -> String {
    let step_yamls = serialize_step_recursive(&root_step);
    step_yamls.join("---\n")
}

pub fn write_step_tree_and_steps_to_file(root_step: Step, filename: &str) -> Result<(), Box<dyn std::error::Error>> {
    let mut output = String::new();
    
    // Add step tree with schema comment
    let schema_comment = "# yaml-language-server: $schema=../../../../../crawler-generator/crawler.schema.json\n";
    let step_tree = step_tree_from_root(root_step.clone());
    output.push_str(&format!("{}{}", schema_comment, step_tree));
    
    // Add document separator
    output.push_str("---\n");
    
    // Add serialized steps (each already has schema comment)
    let serialized_steps = serialize_steps(root_step);
    output.push_str(&serialized_steps);
    
    // Write to file
    fs::write(filename, output)?;
    Ok(())
}

pub fn step_tree_from_root(root_step: Step) -> String {
    // Derive the id from the channel step name
    let id = match &root_step {
        Step::Channel(channel) => {
            let mut name = channel.id.clone();
            // Remove "ListResult" suffix if present
            if name.ends_with("ListResult") {
                name = name.strip_suffix("ListResult").unwrap_or(&name).to_string();
            }
            // Add "Crawler" suffix
            format!("{}Crawler", name)
        }
        _ => "UnknownCrawler".to_string(),
    };

    let yaml_output = YamlOutput {
        kind: "StepTree".to_string(),
        id,
        executor: Executor {
            default: serde_yaml::Value::Mapping(serde_yaml::Mapping::new()),
        },
        root_step: StepNode::from(&root_step),
    };
    
    serde_yaml::to_string(&yaml_output).unwrap_or_else(|e| {
        format!("Error serializing to YAML: {}", e)
    })
}

pub fn op_channel_step(op: &Op, sdk_version: &str) -> Step {
    let mut args = vec![];
    for param in &op.params {
        match param.0.as_str() {
            // ignore API version
            "api-version" => {}
            _ => args.push(param.0.clone()),
        }
    }
    Step::Channel(ChannelStep {
        id: op.response_type.to_string(),
        client_method: ClientMethod {
            package: format!(
                "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/{}/arm{}/{}",
                op.resource_provider, op.resource_provider, sdk_version,
            ),
            // however, client needs to come from the go specs
            client: op.client.to_string(),
            method: op.method.to_string(),
            args,
        },
        children: vec![],
    })
}

pub fn op_resource_step(op: &Op) -> Step {
    Step::Resource(ResourceStep {
        id: op.response_type.to_string(),
        resource: op.response_type.to_string(),
    })
}

pub fn op_payload_step(op: &Op) -> Step {
    Step::Payload(PayloadStep {
        id: op.response_type.to_string(),
        exclude_tags: vec![],
    })
}

pub fn generate_steps(ops: &[Op], sdk_version: &str) -> Step {
    assert!(
        !ops.is_empty(),
        "must have at least one Op to generate steps"
    );

    let mut root_channel = match op_channel_step(&ops[0], sdk_version) {
        Step::Channel(c) => c,
        _ => unreachable!("op_channel_step always returns Step::Channel"),
    };
    root_channel.children.push(op_resource_step(&ops[0]));
    root_channel.children.push(op_payload_step(&ops[0]));

    let mut current = &mut root_channel;

    for op in &ops[1..] {
        let mut next_channel = match op_channel_step(op, sdk_version) {
            Step::Channel(c) => c,
            _ => unreachable!(),
        };

        next_channel.children.push(op_resource_step(op));
        next_channel.children.push(op_payload_step(op));

        current.children.push(Step::Channel(next_channel));

        if let Step::Channel(last_child) = current.children.last_mut().unwrap() {
            current = last_child;
        } else {
            unreachable!("we just pushed a ChannelStep, so this must succeed");
        }
    }

    Step::Channel(root_channel)
}
