package main

import(
	"net"
	"os"
	"os/signal"
	"fmt"
	"log"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/Himanshuxone/gateway/proto"
	"google.golang.org/grpc"
	"database/sql"
	"context"
)

// Database will contain database persistent connection
type Database struct{
	db *sql.DB
}

func main(){
	database := &Database{}
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Errorf("%q",err)
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, database)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx := context.Background()
	go func(ctx context.Context) {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}(ctx)

	// start gRPC server
	log.Println("starting gRPC server...")
	server.Serve(listen)
}

// Create grpc call will create a user
func (db *Database) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error){
	fmt.Printf("\n%v",req)
	return &v1.CreateUserResponse{}, nil
}

// ReadAll will read all the users
func (db *Database) ReadAll(ctx context.Context, req *empty.Empty) (*v1.ReadAllUserResponse, error){
	fmt.Printf("\n%v",req)
	return &v1.ReadAllUserResponse{}, nil
}

