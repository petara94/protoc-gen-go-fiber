package main

import (
	context "context"

	"example/gen/go/greeter"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/zerolog/log"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	greeterpb.UnimplementedGreeterServiceServer
}

func (s *service) SayHello(ctx context.Context, request *greeterpb.HelloRequest) (*greeterpb.HelloResponse, error) {
	return &greeterpb.HelloResponse{
		Message: "Hello " + gofakeit.Name(),
	}, nil
}

func (s *service) PrintRandomImagePNGPathParse(ctx context.Context, request *greeterpb.PrintRandomImagePNGRequest) (*httpbody.HttpBody, error) {
	img := gofakeit.ImagePng(int(request.GetX()), int(request.GetY()))

	return &httpbody.HttpBody{
		ContentType: "image/png",
		Data:        img,
		Extensions:  nil,
	}, nil
}

func (s *service) PrintRandomImagePNGQueryParse(ctx context.Context, request *greeterpb.PrintRandomImagePNGRequest) (*httpbody.HttpBody, error) {
	img := gofakeit.ImagePng(int(request.GetX()), int(request.GetY()))

	return &httpbody.HttpBody{
		ContentType: "image/png",
		Data:        img,
		Extensions:  nil,
	}, nil
}

func (s *service) TestTypesRead(ctx context.Context, request *greeterpb.TestTypesReadRequest) (*emptypb.Empty, error) {
	log.Info().Any("request", request).Msg("TestTypesRead")

	return &emptypb.Empty{}, nil
}
