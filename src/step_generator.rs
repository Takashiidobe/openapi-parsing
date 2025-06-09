// Function to take Vec<Op> -> Step

use crate::openapi_parser::Op;

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash)]
pub struct ClientMethod {
    package: String,
    client: String,
    method: String,
    args: Vec<String>,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash)]
pub struct ChannelStep {
    id: String,
    client_method: ClientMethod,
    children: Vec<Step>,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash)]
pub struct ResourceStep {
    id: String,
    resource: String,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash)]
pub struct PayloadStep {
    id: String,
    exclude_tags: Vec<String>,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash)]
pub struct StepTree {
    root_step: Step,
}

#[derive(Debug, Clone, PartialEq, PartialOrd, Ord, Eq, Hash)]
pub enum Step {
    Channel(ChannelStep),
    Resource(ResourceStep),
    Payload(PayloadStep),
}

pub fn op_channel_step(op: &Op) -> Step {
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
            package: "armcosmos/v3".to_string(),
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

pub fn generate_steps(ops: &[Op]) -> Step {
    assert!(
        !ops.is_empty(),
        "must have at least one Op to generate steps"
    );

    let mut root_channel = match op_channel_step(&ops[0]) {
        Step::Channel(c) => c,
        _ => unreachable!("op_channel_step always returns Step::Channel"),
    };
    root_channel.children.push(op_resource_step(&ops[0]));
    root_channel.children.push(op_payload_step(&ops[0]));

    let mut current = &mut root_channel;

    for op in &ops[1..] {
        let mut next_channel = match op_channel_step(op) {
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
