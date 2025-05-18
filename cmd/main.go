package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"godoallserver/api"
	pb "godoallserver/proto"

	"google.golang.org/grpc"
)

func main() {
	//REST and GraphQL on one HTTP server

	go func() {
		mux := http.NewServeMux()
		api.RegisterREST(mux)
		api.RegisterGraphQL(mux)

		fmt.Println("Starting REST/GraphQL server on :8080")
		log.Fatal(http > ListenAndServer(":8080", mux))
	}()

	//gRPC server on separate port

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMyServiceServer(grpcServer, &api.GRPCServer{})

	fmt.Println("Starting gRPC server on :9090")
	log.Fatal(grpcServer.Serve(lis))
}
