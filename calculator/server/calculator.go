package main

import (
	"context"
	"log"

	"github.com/atselvan/go-grpc/calculator/proto"
)

func (s *Server) Add(ctx context.Context, in *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	log.Printf("Adding (%d) and (%d)", in.Num1, in.Num2)
	return &proto.CalculatorResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}

func (s *Server) Sub(ctx context.Context, in *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	log.Printf("Subtracting (%d) and (%d)", in.Num1, in.Num2)
	return &proto.CalculatorResponse{
		Result: in.Num1 - in.Num2,
	}, nil
}

func (s *Server) Mul(ctx context.Context, in *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	log.Printf("Multiplying (%d) and (%d)", in.Num1, in.Num2)
	return &proto.CalculatorResponse{
		Result: in.Num1 * in.Num2,
	}, nil
}

func (s *Server) Div(ctx context.Context, in *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	log.Printf("Dividing (%d) and (%d)", in.Num1, in.Num2)
	return &proto.CalculatorResponse{
		Result: in.Num1 / in.Num2,
	}, nil
}
