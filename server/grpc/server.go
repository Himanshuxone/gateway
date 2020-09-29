package grpc

import(
	"net"
	"os"
	"os/signal"
	"fmt"
	"log"
	_ "github.com/google/uuid"
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

func RunServer(ctx context.Context, port string){
	log.Println(port)
	database := &Database{}
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		errStr := fmt.Errorf("%s",err)
		log.Panic(errStr)
	}

	// register service
	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, database)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

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
func (db *Database) Create(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error){
	fmt.Printf("\nUser request Object:=> %+v",req)
	return &proto.CreateUserResponse{
		Response: &proto.UserResponse{
			Id: int64(12345),
			User: req.User.GetUser(),
		},
	}, nil
}

// ReadAll will read all the users
func (db *Database) ReadAll(ctx context.Context, req *empty.Empty) (*proto.ReadAllUserResponse, error){
	fmt.Printf("\n%v",req)
	return &proto.ReadAllUserResponse{}, nil
}

