package main

import (
	"fmt"

	"github.com/xakepp35/pkg/xerrors"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func generateFile(plugin *protogen.Plugin) error {
	flagInit()
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, f := range plugin.Files {
		if !f.Generate || len(f.Services) == 0 {
			continue
		}

		g := plugin.NewGeneratedFile(f.GeneratedFilenamePrefix+"_fiber.gw.go", f.GoImportPath)

		genHeader(plugin, g, f)

		g.P("package ", f.GoPackageName)

		genImports(g)

		for _, service := range f.Services {
			err := genService(g, service)
			if err != nil {
				return xerrors.Err(err).Msg("service generation").Str("service", service.GoName).Err()
			}
		}
	}

	return nil
}

func genImports(g *protogen.GeneratedFile) {
	g.Import(contextImport)
	g.Import(stringsImport)
	g.Import(fiberImport)
	g.Import(grpcImport)
	g.Import(grpcMetadataImport)
	g.Import(errorHandlersImport)
	g.Import(errorsBuilderImport)
	g.Import(protoCodesImport)
	g.Import(parsersImport)
}

func genHeader(plugin *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File) {
	g.P("// Code generated by protoc-gen-go-fiber. DO NOT EDIT.")
	g.P("// versions:")
	protocVersion := "(unknown)"
	if v := plugin.Request.GetCompilerVersion(); v != nil {
		protocVersion = fmt.Sprintf("v%v.%v.%v", v.GetMajor(), v.GetMinor(), v.GetPatch())
		if s := v.GetSuffix(); s != "" {
			protocVersion += "-" + s
		}
	}
	g.P("// \tprotoc-gen-go-fiber ", version)
	g.P("// \tprotoc              ", protocVersion)
	g.P("// source: ", f.Desc.Path())
	g.P()
}
