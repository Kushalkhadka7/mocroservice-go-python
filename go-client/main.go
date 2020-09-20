package main

import (
	"client/authclient"
	"client/laptopclient"
	pb "client/pb"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	authConn := createAuthConn()
	authClient := authclient.NewAuthClient(authConn, "kushal", "kushal")
	interceptor, err := authclient.NewAuthInterceptor(authClient, 30*time.Second)
	if err != nil {
		panic(err)
	}

	serverConn, err := grpc.Dial("localhost:8080",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		panic(err)
	}

	laptopClient := laptopclient.NewLapotpClient(serverConn)

	laptopClient.CreateNewLaptop()

}

func createUser(ctx context.Context, authClient pb.AuthServiceClient) {

	req := &pb.CreateUserRequest{
		User: &pb.User{
			Name:     "kushal",
			Password: "kushal",
			Role:     "admin",
		},
	}

	res, err := authClient.CreateUser(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println("Successfully created user")
}

func createAuthConn() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}
