//go:build ignore

package main

import (
	"log"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{
		Hooks: []gen.Hook{
			addBoolOmitempty(),
		},
		Features: []gen.Feature{
			gen.FeatureExecQuery,
		},
	}); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func addBoolOmitempty() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, f := range node.Fields {
					if gen.Field.IsBool(*f) {
						f.StructTag = strings.Replace(f.StructTag, ",omitempty", "", 1)
						f.StructTag = strings.Replace(f.StructTag, "omitempty", "", 1)
					}
				}
			}
			return next.Generate(g)
		})
	}
}
