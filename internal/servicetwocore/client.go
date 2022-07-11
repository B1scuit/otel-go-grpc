package servicetwocore

import "github.com/B1scuit/otel-go-grpc/proto"

type ClientOptions struct{}

type Client struct {
	proto.UnimplementedServiceTwoServer
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
