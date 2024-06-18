package main

import (
	mygrpc "github.com/Ilhamkawe/gRPC-unary-server/internal/adapter/grpc"
	"github.com/Ilhamkawe/gRPC-unary-server/internal/application"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &application.HelloService{}

	grpcAdapter := mygrpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}
