package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ltbatista/prime-number-decomposition/prime/primepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Prime(req *primepb.PrimeRequest, stream primepb.PrimeService_PrimeServer) error {
	inteiro := req.GetPrime().GetNumero()
	fmt.Printf("Initializing prime decomposition for number %v...\n", inteiro)
	k := 2
	for n := int(inteiro); n > 1; {
		if n%k == 0 {
			result := strconv.Itoa(k)
			res := &primepb.PrimeResponse{
				Result: result,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			n = n / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello, I'm the prime server.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	primepb.RegisterPrimeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
