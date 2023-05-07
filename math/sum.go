package math

import (
	"context"
	"http2/proto"
)

type GRPCSum struct {}

func (s *GRPCSum) Sum(ctx context.Context, req *proto.MathRequest) (*proto.MathResponse, error) {
	return &proto.MathResponse{Result: req.X + req.Y}, nil
}