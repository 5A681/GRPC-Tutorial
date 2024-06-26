package services

import (
	"context"
	"fmt"
	"io"
	"time"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
	Average(numbers ...float64) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Printf("Service: Hello\n")
	fmt.Printf("Request: %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)
	return nil
}
func (base calculatorService) Average(numbers ...float64) error {

	stream, err := base.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Service: Fibonacci\n")
	for _, number := range numbers {
		req := AverageRequest{
			Number: number,
		}
		stream.Send(&req)
		fmt.Printf("Request: %v\n", req.Number)
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("Response: %v\n", res.Result)
	return nil
}

func (base calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := base.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}
	fmt.Printf("Service: Fibonacci\n")
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Request: %v\n", req.N)
		fmt.Printf("Response: %v\n", res.Result)
	}

	return nil
}
