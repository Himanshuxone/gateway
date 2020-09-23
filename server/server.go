package main

import(
	grpcServer "github.com/Himanshuxone/gateway/grpc"
	restServer "github.com/Himanshuxone/gateway/rest"
)

const (
	GRPCPort = "9090"
	HTTPPort = "9000"
)

func main(){

	ctx := context.Background()
	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, GRPCPort, HTTPPort)
	}()

	return grpc.RunServer()
}