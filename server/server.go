package main

import(
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/Himanshuxone/gateway/proto/v1"
	"google.golang.org/grpc"
	"database/sql"
)

// Database will contain database persistent connection
type Database struct{
	db *sql.DB
}

func main(){
	database := Database{}
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, database)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}

// Create grpc call will create a user
func (db *Database) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error){
	fmt.Printf("\n%#v",req)
}

// ReadAll will read all the users
func (db *Database) ReadAll(ctx context.Context, req *empty.Empty) (*ReadAllUserResponse, error){
	fmt.Printf("\n%#v",req)
}

