package main

import (
	"context"
	grpcServer "github.com/Himanshuxone/gateway/server/grpc"
	restServer "github.com/Himanshuxone/gateway/server/rest"
)

var (
	GRPCPort = "9090"
	HTTPPort = "9000"
)

func main() {

	ctx := context.Background()

	// run HTTP gateway
	go func() {
		_ = restServer.RunServer(ctx, GRPCPort, HTTPPort)
	}()

	grpcServer.RunServer(ctx, GRPCPort)
}
