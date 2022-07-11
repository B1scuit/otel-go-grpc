package grpcserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServerInterface interface {
	RegisterService(*grpc.ServiceDesc, interface{})
}

type CoreInterface interface{}

type ClientOptions struct {
	Port int

	ServiceDesc *grpc.ServiceDesc
	Core        interface{}
}

type Client struct {
	listener net.Listener
	srv      *grpc.Server
}

func New(opts *ClientOptions) (*Client, error) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", opts.Port))
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()

	srv.RegisterService(opts.ServiceDesc, opts.Core)

	reflection.Register(srv)

	return &Client{
		listener: listener,
		srv:      srv,
	}, nil
}

func Must(client *Client, err error) *Client {
	if err != nil {
		panic(err)
	}

	return client
}

func (c *Client) Run() error {
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop

		c.srv.GracefulStop()
	}()

	return c.srv.Serve(c.listener)
}
