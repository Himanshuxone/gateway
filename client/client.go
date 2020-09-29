package main

import (
	"log"
	"google.golang.org/grpc"
	"context"
	"time"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/Himanshuxone/gateway/proto"
)

func main(){
	// get configuration
	address := "localhost:9090"

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("connection state ====> ", conn.GetState())

	// user client
	Client(conn)
}

// Client is a user service client
func Client(conn *grpc.ClientConn){

	c := proto.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	req1 := proto.CreateUserRequest{
		User: &proto.User{
			User: "Irene A. Andrews",
			Email: "IreneAAndrews@teleworm.us",
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	user := res1.Response.GetUser()
	log.Println("User inserted: ", user)

	res4, err := c.ReadAll(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res4)

}