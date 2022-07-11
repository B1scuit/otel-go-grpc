package serviceonecore

import "github.com/B1scuit/otel-go-grpc/proto"

type ClientOptions struct{}

type Client struct {
	proto.UnimplementedServiceOneServer
}

func New(opts *ClientOptions) (*Client, error) {
	return &Client{}, nil
}

func Must(client *Client, err error) *Client {
	if err != nil {
		panic(err)
	}

	return client
}
