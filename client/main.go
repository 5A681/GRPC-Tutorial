package main

import (
	"client/services"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:50051", "gRPC server host")
	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()
	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}

	} else {
		creds = insecure.NewCredentials()
	}
	cc, err = grpc.NewClient(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)
	err = calculatorService.Average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if err != nil {
		log.Fatal(err)
	}
}
