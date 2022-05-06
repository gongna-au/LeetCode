package main

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/generator/descriptor"
)

type netrpcPlugin struct {
	*generator.Generator
}

func (p *netrpcPlugin) Name() string {
	return "netRpc"
}
func (p *netrpcPlugin) Init(g generator.Generator) {

	p.Generator = g
}
func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)

	}

}
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}

}
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	fmt.Print("// TODO: import code")

}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	fmt.Print("// TODO: service code, Name = " + svc.GetName())
}
func init() {
	generator.RegisterPlugin(new(netrpcPlugin))
}
