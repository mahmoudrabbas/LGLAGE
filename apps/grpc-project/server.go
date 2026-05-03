package main

import (
	"fmt"
	"grpc-demo/db"
	pb "grpc-demo/proto"
	"grpc-demo/repository"
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

	repo := repository.NewUserRepository(db.DB)

	userService := service.NewUserService(repo)

	pb.RegisterUserServiceServer(grpcServer, userService)

	fmt.Println("gRPC Server running on port 50051 🚀")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
