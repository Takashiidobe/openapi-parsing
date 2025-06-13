package goparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type MethodInfo struct {
	Name    string
	Returns []string
}

func exprString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return "*" + exprString(e.X)
	case *ast.SelectorExpr:
		return exprString(e.X) + "." + e.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprString(e.Elt)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", exprString(e.Key), exprString(e.Value))
	default:
		return fmt.Sprintf("%T", expr)
	}
}

func Parse(dir string) (map[string][]MethodInfo, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]MethodInfo)
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				switch d := decl.(type) {
				case *ast.FuncDecl:
					if d.Recv != nil && len(d.Recv.List) > 0 {
						var recv string
						switch t := d.Recv.List[0].Type.(type) {
						case *ast.StarExpr:
							if id, ok := t.X.(*ast.Ident); ok {
								recv = id.Name
							}
						case *ast.Ident:
							recv = t.Name
						}
						if recv != "" {
							var rets []string
							if d.Type.Results != nil {
								for _, res := range d.Type.Results.List {
									rets = append(rets, exprString(res.Type))
								}
							}
							result[recv] = append(result[recv], MethodInfo{Name: d.Name.Name, Returns: rets})
						}
					}
				}
			}
		}
	}
	return result, nil
}
