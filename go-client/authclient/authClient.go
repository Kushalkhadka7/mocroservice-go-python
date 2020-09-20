package authclient

import (
	pb "client/pb"
	"time"

	"context"

	"google.golang.org/grpc"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(conn *grpc.ClientConn, username, password string) *AuthClient {
	service := pb.NewAuthServiceClient(conn)
	return &AuthClient{
		service, username, password,
	}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := pb.CreateUserRequest{
		User: &pb.User{
			Name:     "kushal",
			Password: "kushal",
			Role:     "admin",
		},
	}

	res, err := client.service.CreateUser(ctx, &req)
	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}
