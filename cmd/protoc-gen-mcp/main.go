package main

import (
	"flag"

	"github.com/omrikiei/protoc-gen-mcp/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := generator.GenerateFile(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}
