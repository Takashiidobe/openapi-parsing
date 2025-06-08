use anyhow::Result;
use gosyn::{
    ast::{Declaration, Expression, FuncDecl, Ident, StructType, TypeSpec},
    parse_dir,
};
use std::{collections::HashMap, path::Path};

#[derive(Debug)]
pub struct StructInfo {
    /// The AST node for the struct’s definition
    def: StructType,
    /// All methods whose receiver type matches this struct
    methods: Vec<FuncDecl>,
}

#[derive(Debug)]
pub struct MethodInfo {
    /// the method’s identifier
    name: String,
    /// return types (just types; Go allows unnamed results)
    returns: Vec<String>,
}

fn expr_to_string(expr: &Expression) -> String {
    match expr {
        Expression::Ident(id) => id.name.clone(),
        Expression::TypePointer(ptr) => {
            let inner = expr_to_string(&ptr.typ);
            format!("*{}", inner)
        }
        // slice, e.g. "[]byte"
        Expression::TypeSlice(slice) => {
            // Depending on your gosyn version the field may be named `elem` or `typ`
            let inner = expr_to_string(&slice.typ);
            format!("[]{}", inner)
        }

        Expression::Index(idx) => {
            let container = expr_to_string(&idx.left);
            let index = expr_to_string(&idx.index);
            format!("{}[{}]", container, index)
        }

        // map, e.g. "map[string]int"
        Expression::TypeMap(m) => {
            let key = expr_to_string(&m.key);
            let val = expr_to_string(&m.val);
            format!("map[{}]{}", key, val)
        }
        Expression::Selector(sel) => {
            let pkg = expr_to_string(&sel.x);
            let field = sel.sel.name.clone();
            format!("{}.{}", pkg, field)
        }

        // … handle other cases
        _ => format!("{:?}", expr), // fallback to debug
    }
}

pub fn parse(path: &str) -> Result<HashMap<String, Vec<MethodInfo>>> {
    // 1) Parse the directory into a map: package_path → Package
    let pkgs = parse_dir(path)?; // :contentReference[oaicite:0]{index=0}

    // 2) Pick the first (and typically only) package
    let (_pkg_path, pkg) = pkgs
        .into_iter()
        .next()
        .expect("no Go package found in directory");

    // 3) Build a map: struct name → StructInfo
    let mut structs: HashMap<String, StructInfo> = HashMap::new();

    // --- First pass: collect all struct type specs ---
    for file in &pkg.files {
        for decl in &file.decl {
            if let Declaration::Type(typed) = decl {
                // :contentReference[oaicite:1]{index=1}
                for spec in &typed.specs {
                    let TypeSpec { name, typ, .. } = spec;
                    // Look for `type Foo struct { ... }`
                    if let Expression::TypeStruct(st) = &typ {
                        // :contentReference[oaicite:2]{index=2}
                        structs.insert(
                            name.name.clone(),
                            StructInfo {
                                def: st.clone(),
                                methods: Vec::new(),
                            },
                        );
                    }
                }
            }
        }
    }

    // --- Second pass: collect all methods and attach to their receiver’s struct ---
    for file in &pkg.files {
        for decl in &file.decl {
            if let Declaration::Function(fn_decl) = decl {
                if let Some(recv_list) = &fn_decl.recv {
                    // assume the first field in recv_list is the receiver
                    if let Some(first) = recv_list.list.first() {
                        // extract the receiver’s type identifier
                        let recv_ty = match &first.typ {
                            // pointer receiver `*MyType`
                            Expression::TypePointer(ptr) => match &*ptr.typ {
                                Expression::Ident(id) => id.name.clone(), // ← here
                                _ => continue,
                            },
                            // non-pointer receiver `MyType`
                            Expression::Ident(id) => id.name.clone(), // ← and here
                            _ => continue,
                        };

                        // push this FuncDecl into the matching StructInfo
                        structs
                            .entry(recv_ty)
                            .or_insert_with(|| StructInfo {
                                def: StructType {
                                    pos: (0, 0),
                                    fields: Vec::new(),
                                },
                                methods: Vec::new(),
                            })
                            .methods
                            .push(fn_decl.clone());
                    }
                }
            }
        }
    }
    let mut methods_by_struct: HashMap<String, Vec<MethodInfo>> = HashMap::new();

    for (struct_name, info) in &structs {
        let infos = info
            .methods
            .iter()
            .map(|m: &FuncDecl| {
                // Method name
                let name = m.name.name.clone();

                // Collect return types (unnamed or named both appear in the list)
                let returns = m
                    .clone()
                    .typ
                    .result
                    .list
                    .into_iter()
                    .map(|x| expr_to_string(&x.typ))
                    .collect();

                MethodInfo { name, returns }
            })
            .collect();

        methods_by_struct.insert(struct_name.clone(), infos);
    }

    Ok(methods_by_struct)
}
