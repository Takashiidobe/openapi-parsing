package openapi

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Op struct {
	Path        string
	Method      string
	OperationID string
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
	var ops []Op
	for pair := p.model.Model.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
		path := pair.Key()
		item := pair.Value()
		if item.Get != nil {
			op := Op{Path: path, Method: "GET", OperationID: item.Get.OperationId}
			ops = append(ops, op)
		}
	}
	return ops
}
