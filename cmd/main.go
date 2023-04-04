package main

import (
	"context"
	"grpc-demo/greeter"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myGreeterServer struct {
	greeter.UnimplementedGreeterServer
}

func (server myGreeterServer) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return nil, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greeterServer := &myGreeterServer{}
	greeter.RegisterGreeterServer(s, greeterServer)
	log.Println("server listening at :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
