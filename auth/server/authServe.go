package server

import (
	auth "auth/pbauth"
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
func (server *AuthServer) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	fmt.Println("i am called")
	fmt.Println(req.User)
	res, err := server.UserService.SaveUser(ctx, req.User)
	if err != nil {
		fmt.Println("i am called1")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)

	return &auth.CreateUserResponse{
		Success: "Success",
	}, nil
}
