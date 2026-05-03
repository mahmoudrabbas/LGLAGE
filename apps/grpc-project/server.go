package main

import (
	"fmt"
	"grpc-demo/db"
	pb "grpc-demo/proto"
	"grpc-demo/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	db.Connect()

	if db.DB == nil {
		log.Fatal("DB is nil")
	}

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	userService := service.NewUserService(db.DB)

	pb.RegisterUserServiceServer(grpcServer, userService)

	fmt.Println("gRPC Server running on port 50051 🚀")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
