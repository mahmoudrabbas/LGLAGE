package service

import (
	"context"
	"database/sql"

	// "grpc-demo/db"
	pb "grpc-demo/proto"
	"grpc-demo/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	repo *repository.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	id, name, email, err := s.repo.CreateUser(req.Name, req.Email)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.UserResponse{
		User: &pb.User{
			Id:    id,
			Name:  name,
			Email: email,
		},
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {

	userId, name, email, err := s.repo.GetUserById(req.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.UserResponse{
		User: &pb.User{
			Id:    userId,
			Name:  name,
			Email: email,
		},
	}, nil

}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	id, name, email, err := s.repo.UpdateUser(req.Id, req.Name, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "internal Error")

	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:    id,
			Name:  name,
			Email: email,
		},
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	row, err := s.repo.DeleteUser(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())

	}
	if row == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
