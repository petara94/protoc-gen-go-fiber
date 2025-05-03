package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	flags                         = flag.NewFlagSet("protoc-gen-go-fiber", flag.ExitOnError)
	flagErrorHandlersPackage      = flags.String("error_handlers_package", "github.com/petara94/protoc-gen-go-fiber/utils", "package with error handlers funcs")
	flagJsonUnmarshalPackage      = flags.String("json_unmarshal_package", "encoding/json", "package with json unmarshalers")
	flagGrpcErrorHandleFunc       = flags.String("grpc_error_handle_func", "HandleGRPCStatusError", "func name for handle grpc error")
	flagUnmarshalErrorHandleFunc  = flags.String("unmarshal_error_handle_func", "HandleUnmarshalError", "func name for handle unmarshal error")
	flagValidationErrorHandleFunc = flags.String("validation_error_handle_func", "HandleValidationError", "func name for handle validation error")
	flagDisableGrpcInterceptor    = flags.Bool("disable_grpc_interceptor", false, "disable grpc interceptor")
)

func flagInit() {
	errorHandlersImport = protogen.GoImportPath(*flagErrorHandlersPackage)
	jsonUnmarshalImport = protogen.GoImportPath(*flagJsonUnmarshalPackage)
}
