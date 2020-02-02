package main

import (
	"context"
	"log"
	"net"

	pb "github.com/zuiurs/grpc-web-sample/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type GreeterServer struct{}

func (g *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Receive: %s", req.GetName())
	return &pb.HelloReply{Message: req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreeterServer{})
	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start greeter server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
