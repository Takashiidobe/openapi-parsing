use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};
use walkdir::WalkDir;

fn find_json_spec(api_ver: &str, target: &str) -> io::Result<Vec<String>> {
    let version_segment = format!("/stable/{}/", api_ver);
    let mut matches = Vec::new();

    for entry in WalkDir::new("../azure-rest-api-specs/specification")
        .into_iter()
        .filter_map(Result::ok)
    {
        let path = entry.path();
        let path_s = path.to_string_lossy();

        if !path_s.contains(&version_segment)
            || path_s.contains("/preview/")
            || path.extension().and_then(|e| e.to_str()) != Some("json")
        {
            continue;
        }

        let file = File::open(path)?;
        let reader = BufReader::new(file);
        if reader.lines().any(|line| match line {
            Ok(l) => l.contains(target),
            Err(_) => false,
        }) {
            matches.push(path.to_string_lossy().to_string());
        }
    }

    Ok(matches)
}

pub fn spec_finder(api_version: &str, target: &str) -> io::Result<Vec<String>> {
    let specs = find_json_spec(api_version, target)?;
    Ok(specs)
}
