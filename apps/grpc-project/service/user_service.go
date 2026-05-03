package service

import (
	"context"
	"database/sql"
	"grpc-demo/db"
	pb "grpc-demo/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	DB *sql.DB
	pb.UnimplementedUserServiceServer
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	result, err := s.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", req.Name, req.Email)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:    id,
			Name:  req.Name,
			Email: req.Email,
		},
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {

	row := db.DB.QueryRow("SELECT id, name, email FROM users WHERE id=?", req.Id)

	var user pb.User

	err := row.Scan(&user.Id, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &pb.UserResponse{
		User: &user,
	}, nil

}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	var exists int
	err := s.DB.QueryRow(
		"SELECT 1 FROM users WHERE id = ?",
		req.Id,
	).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	_, err = s.DB.Exec(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
		req.Name,
		req.Email,
		req.Id,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:    req.Id,
			Name:  req.Name,
			Email: req.Email,
		},
	}, nil

}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	res, err := s.DB.Exec("DELETE FROM users where id=?", req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	row, err := res.RowsAffected()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if row == 0 {
		return nil, status.Errorf(codes.NotFound, "User Not found")
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
