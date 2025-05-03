package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func generateRouteMethodName(method *protogen.Method) string {
	return fmt.Sprintf("__%s_%s_Route", method.Parent.GoName, method.GoName)
}

func generateFiberMethodRote(g *protogen.GeneratedFile, method *protogen.Method) {
	opts := method.Desc.Options().(*descriptorpb.MethodOptions)

	methodType, httpPath := grpcOptionToMethodAndPathString(opts)
	if httpPath == "/" {
		httpPath += string(method.Parent.Desc.FullName()) + "/" + string(method.Desc.Name())
	}

	g.P("	app.", methodType, `("`, httpPath, `", `, generateRouteMethodName(method), `)`)
}
