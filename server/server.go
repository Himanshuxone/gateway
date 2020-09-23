package main

import (
	"github.com/Himanshuxone/gateway/proto"
	grpcServer "github.com/Himanshuxone/gateway/server/grpc"
	restServer "github.com/Himanshuxone/gateway/server/rest"
)

const (
	GRPCPort = "9090"
	HTTPPort = "9000"
)

func main() {

	ctx := context.Background()

	// run HTTP gateway
	go func() {
		_ = restServer.RunServer(ctx, GRPCPort, HTTPPort)
	}()

	return grpcServer.RunServer(ctx, GRPCPort)
}
