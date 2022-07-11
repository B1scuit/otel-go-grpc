package main

import (
	"log"

	"github.com/B1scuit/otel-go-grpc/internal/servicethreecore"
	"github.com/B1scuit/otel-go-grpc/pkg/grpcserver"
	"github.com/B1scuit/otel-go-grpc/proto"
)

func main() {
	grpcServer := grpcserver.Must(grpcserver.New(&grpcserver.ClientOptions{
		Port:        8080,
		ServiceDesc: &proto.ServiceThree_ServiceDesc,
		Core:        servicethreecore.Must(servicethreecore.New(&servicethreecore.ClientOptions{})),
	}))

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
