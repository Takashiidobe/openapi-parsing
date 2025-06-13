package openapi

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Parameter struct {
	Name string
	Type string
}

type Op struct {
	Path             string
	Params           []Parameter
	Client           string
	Method           string
	ResourceProvider string
	ResponseType     string
}

type Parser struct {
	model *libopenapi.DocumentModel[v3.Document]
}

func NewParser(path string) (*Parser, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	doc, err := libopenapi.NewDocument(data)
	if err != nil {
		return nil, err
	}
	model, errs := doc.BuildV3Model()
	if len(errs) > 0 {
		return nil, fmt.Errorf("error building model: %v", errs[0])
	}
	return &Parser{model: model}, nil
}

func (p *Parser) Ops() []Op {
	if p == nil || p.model == nil {
		return nil
	}

	pathItems := p.model.Model.Paths.PathItems
	var ops []Op

	for pair := pathItems.First(); pair != nil; pair = pair.Next() {
		key := pair.Key()
		item := pair.Value()
		if !strings.HasSuffix(key, ".get") || item.Get == nil {
			continue
		}

		baseKey := strings.TrimSuffix(key, ".get")
		baseItem := pathItems.GetOrZero(baseKey)
		if baseItem == nil {
			continue
		}

		var meta struct {
			Path string `yaml:"path"`
		}
		if ext, ok := baseItem.Extensions.Get("x-ms-metadata"); ok {
			_ = ext.Decode(&meta)
		}

		op := item.Get
		opID := op.OperationId
		split := strings.Split(opID, "_")
		if len(split) < 2 {
			continue
		}

		var resourceProvider string
		for _, seg := range strings.Split(meta.Path, "/") {
			if strings.HasPrefix(seg, "Microsoft.") {
				resourceProvider = strings.TrimPrefix(seg, "Microsoft.")
				break
			}
		}

		client := strings.TrimPrefix(split[0]+"Client", resourceProvider)
		method := "New" + split[1] + "Pager"

		var params []Parameter
		if op.Parameters != nil {
			for _, param := range op.Parameters {
				var typ string
				if param.Schema != nil {
					schema := param.Schema.Schema()
					if schema != nil {
						typ = strings.Join(schema.Type, ",")
					}
				}
				params = append(params, Parameter{Name: param.Name, Type: typ})
			}
		}

		var respType string
		if op.Responses != nil {
			resp := op.Responses.FindResponseByCode(200)
			if resp != nil && resp.Content != nil {
				mt := resp.Content.GetOrZero("application/json")
				if mt != nil && mt.Schema != nil {
					sch := mt.Schema.Schema()
					if sch != nil && sch.Extensions != nil {
						if n, ok := sch.Extensions.Get("x-ms-metadata"); ok {
							var r struct {
								Name string `yaml:"name"`
							}
							if err := n.Decode(&r); err == nil {
								respType = r.Name
							}
						}
					}
				}
			}
		}

		ops = append(ops, Op{
			Path:             meta.Path,
			Params:           params,
			Client:           client,
			Method:           method,
			ResourceProvider: strings.ToLower(resourceProvider),
			ResponseType:     respType,
		})
	}

	sort.Slice(ops, func(i, j int) bool {
		if ops[i].Path == ops[j].Path {
			return ops[i].ResponseType < ops[j].ResponseType
		}
		return ops[i].Path < ops[j].Path
	})

	return ops
}

func prefixes(path string) []string {
	trimmed := strings.TrimPrefix(path, "/")
	var acc string
	var out []string
	for _, seg := range strings.Split(trimmed, "/") {
		acc += "/" + seg
		out = append(out, acc)
	}
	return out
}

func FindDependencies(ops []Op, targetPath string) []Op {
	prefixSet := map[string]struct{}{}
	for _, p := range prefixes(targetPath) {
		prefixSet[p] = struct{}{}
	}
	var deps []Op
	for _, op := range ops {
		if _, ok := prefixSet[op.Path]; ok {
			deps = append(deps, op)
		}
	}
	sort.Slice(deps, func(i, j int) bool {
		return len(deps[i].Path) < len(deps[j].Path)
	})
	return deps
}
