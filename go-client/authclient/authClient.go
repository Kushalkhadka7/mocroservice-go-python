package authclient

import (
	pb "client/pb"
	"time"

	"context"

	"google.golang.org/grpc"
)

// AuthClient struct.
type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

// NewAuthClient creates new authentication client.
func NewAuthClient(conn *grpc.ClientConn, username, password string) *AuthClient {
	service := pb.NewAuthServiceClient(conn)
	return &AuthClient{
		service, username, password,
	}
}

// Login calls auth modules to login the user and get jwt token.
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := pb.CreateUserRequest{
		User: &pb.User{
			Name:     client.username,
			Password: client.password,
			Role:     "admin",
		},
	}

	res, err := client.service.CreateUser(ctx, &req)
	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}
