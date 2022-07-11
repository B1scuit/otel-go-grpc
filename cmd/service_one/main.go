package main

import (
	"log"

	"github.com/B1scuit/otel-go-grpc/internal/serviceonecore"
	"github.com/B1scuit/otel-go-grpc/pkg/grpcserver"
	"github.com/B1scuit/otel-go-grpc/proto"
)

func main() {
	grpcServer := grpcserver.Must(grpcserver.New(&grpcserver.ClientOptions{
		Port:        8080,
		ServiceDesc: &proto.ServiceOne_ServiceDesc,
		Core:        serviceonecore.Must(serviceonecore.New(&serviceonecore.ClientOptions{})),
	}))

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
