package main

import (
	"fmt"
	"log"

	"openapi-parsing/pkg/openapi"
)

func main() {
	parser, err := openapi.NewParser("json/documentdb/openapi-document.json")
	if err != nil {
		log.Fatalf("openapi parse: %v", err)
	}
	ops := parser.Ops()
	fmt.Printf("loaded %d ops\n", len(ops))
	for i, op := range ops {
		if i > 4 {
			break
		}
		fmt.Printf("%s %s -> %s\n", op.Method, op.Path, op.ResponseType)
	}
}
