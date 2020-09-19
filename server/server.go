package main

import(
	"net/http"
	"fmt"
	"google.golang.org/grpc"
)

func main(){
	server := grpc.NewServer()
	
}

func handle(response http.ResponseWriter, request *http.Request){
	fmt.Println("Welcome tp grpc serve http")
}