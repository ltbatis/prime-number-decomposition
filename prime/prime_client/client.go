package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ltbatista/prime-number-decomposition/prime/primepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, i'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := primepb.NewPrimeServiceClient(cc)

	doPrimeDecompositionRequest(c)
}

func doPrimeDecompositionRequest(c primepb.PrimeServiceClient) {
	//TODO: Implementar
	fmt.Println("Starting Prime Request RPC...")
	req := &primepb.PrimeRequest{
		Prime: &primepb.Prime{
			Numero: 120,
		},
	}
	resStream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// chegamos ao fim do streaming
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
