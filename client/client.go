package main

import (
	"flag"
	"log"
	"google.golang.org/grpc"
	"github.com/Himanshuxone/gateway/client"
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/Himanshuxone/gateway/proto"
)

func main(){
	// get configuration
	address := "9090"

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
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

	c := v1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	req1 := v1.CreateUserRequest{
		User: &v1.User{
			User: "Irene A. Andrews",
			Email: "IreneAAndrews@teleworm.us",
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	id := res1.Id
	log.Println("User inserted: ", id)

	res4, err := c.ReadAll(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res4)

}