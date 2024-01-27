package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/akihiro-fukuchi/cnasb/envoy/pkg/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	"github.com/akihiro-fukuchi/cnasb/envoy/pkg/proto"
)

type EchoServer struct{}

func (s *EchoServer) Echo(ctx context.Context, in *proto.EchoRequest) (*proto.EchoResponse, error) {
	log.Printf("Handling Echo request [%v] with context %v", in, ctx)
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Unable to get hostname %v", err)
		hostname = ""
	}

	if err := grpc.SendHeader(ctx, metadata.Pairs("hostname", hostname)); err != nil {
		log.Printf("Unable to send header %v", err)
	}
	return &proto.EchoResponse{Content: in.Content}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterEchoServer(grpcServer, &EchoServer{})
	grpc_health_v1.RegisterHealthServer(grpcServer, &health.Server{})
	reflection.Register(grpcServer)
	log.Printf("Listening for Echo on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
