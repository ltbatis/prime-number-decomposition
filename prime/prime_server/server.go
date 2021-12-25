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
	fmt.Println("Initializing prime decomposition...")
	inteiro := req.GetPrime().GetNumero()
	// TODO: Implementar a decomposição dos números primos, por enquanto só está fazendo um contador
	// para testar se o server está funcionando
	for i := 1; i < int(inteiro); i++ {
		result := strconv.Itoa(i)
		res := &primepb.PrimeResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
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
