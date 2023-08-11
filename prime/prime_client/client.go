package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/ltbatista/prime-number-decomposition/prime/primepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, i'm a client!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := primepb.NewPrimeServiceClient(cc)

	doPrimeDecompositionRequest(c)
}

func doPrimeDecompositionRequest(c primepb.PrimeServiceClient) {
	number, _ := strconv.Atoi(os.Args[1])
	fmt.Println("Starting Prime Request RPC...")
	req := &primepb.PrimeRequest{
		Prime: &primepb.Prime{
			Numero: int32(number),
		},
	}
	resStream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Prime RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			log.Print("End.")
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from Prime Decomposition: %v", msg.GetResult())
	}
}
