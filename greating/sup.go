package greating

import (
	"context"
	"http2/proto"
)

type GRPCGreating struct {}

func (s *GRPCGreating) Greet(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	return &proto.GreetResponse{Message: "Hello" + " " + req.Name}, nil
}