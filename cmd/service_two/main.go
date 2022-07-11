package main

import (
	"log"

	"github.com/B1scuit/otel-go-grpc/internal/servicetwocore"
	"github.com/B1scuit/otel-go-grpc/pkg/grpcserver"
	"github.com/B1scuit/otel-go-grpc/proto"
)

func main() {
	grpcServer := grpcserver.Must(grpcserver.New(&grpcserver.ClientOptions{
		Port:        8080,
		ServiceDesc: &proto.ServiceTwo_ServiceDesc,
		Core:        servicetwocore.Must(servicetwocore.New(&servicetwocore.ClientOptions{})),
	}))

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
