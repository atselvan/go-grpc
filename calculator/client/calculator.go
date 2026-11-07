package main

import (
	"context"
	"log"

	"github.com/atselvan/go-grpc/calculator/proto"
)

func add(c proto.CalculatorServiceClient, num1, num2 int32) {
	log.Printf("Adding (%d) and (%d)", num1, num2)
	res, err := c.Add(context.Background(), &proto.CalculatorRequest{
		Num1: num1,
		Num2: num2,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}
	log.Printf("Result: %d\n", res.Result)
}

func sub(c proto.CalculatorServiceClient, num1, num2 int32) {
	log.Printf("Subtracting (%d) and (%d)", num1, num2)
	res, err := c.Sub(context.Background(), &proto.CalculatorRequest{
		Num1: num1,
		Num2: num2,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}
	log.Printf("Result: %d\n", res.Result)
}

func mul(c proto.CalculatorServiceClient, num1, num2 int32) {
	log.Printf("Multiplying (%d) and (%d)", num1, num2)
	res, err := c.Mul(context.Background(), &proto.CalculatorRequest{
		Num1: num1,
		Num2: num2,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}
	log.Printf("Result: %d\n", res.Result)
}

func div(c proto.CalculatorServiceClient, num1, num2 int32) {
	log.Printf("Dividing (%d) by (%d)", num1, num2)
	res, err := c.Div(context.Background(), &proto.CalculatorRequest{
		Num1: num1,
		Num2: num2,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}
	log.Printf("Result: %d\n", res.Result)
}
