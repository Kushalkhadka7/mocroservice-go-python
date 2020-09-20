package server

import (
	pb "auth/pb"
	"auth/service"
	"context"
	"fmt"
)

// AuthServer initializes new auth server.
type AuthServer struct {
	UserService service.UserService
}

// NewAuthServer creates new auth server.
func NewAuthServer(userService service.UserService) *AuthServer {
	return &AuthServer{
		UserService: userService,
	}
}

// CreateUser creates new user.
func (server *AuthServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := server.UserService.SaveUser(ctx, req.User)
	if err != nil {
		fmt.Println("i am called1")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)

	return &pb.CreateUserResponse{
		Success: "Success",
	}, nil
}

func (server *AuthServer) VerifyUser(ctx context.Context, req *pb.VerifyUserTokenRequest) (*pb.VerifyUserTokenResponse, error) {
	res, err := server.UserService.VerifyUser(ctx, req)
	if err != nil {
		fmt.Println("i am called1")
		fmt.Println(err)
		return nil, err
	}

	return &pb.VerifyUserTokenResponse{
		Name: res.UserName,
		Role: res.Role,
	}, nil
}
