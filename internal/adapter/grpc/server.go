package grpc

import (
	"fmt"
	"github.com/Ilhamkawe/gRPC-unary-server/internal/port"
	"github.com/Ilhamkawe/protobuf-grpc/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		grpcPort:     grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port :%d : %v\n", a.grpcPort, err)
	}

	log.Printf("Server listening on port :%d ", a.grpcPort)

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on port %d : %v\n", a.grpcPort, err)
	}
}
