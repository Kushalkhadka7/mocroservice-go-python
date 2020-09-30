package main

import (
	"client/authclient"
	"client/laptopclient"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Connect to auth server.
	authConn, err := createAuthConn()
	authClient := authclient.NewAuthClient(authConn, "kushal", "khadka")
	if err != nil {
		panic(err)
	}

	interceptor, err := authclient.NewAuthInterceptor(authClient, 30*time.Second)
	if err != nil {

		fmt.Println(err)
		panic(err)
	}

	// Connect to main server using the jwt token get from authentication.
	serverConn, err := createServerConn(interceptor)
	if err != nil {
		panic(err)
	}

	service := pb.NewLaptopServiceClient(serverConn)
	laptopClient := laptopclient.NewLapotpClient(service)

	// Calls server methods.
	laptopClient.CreateNewLaptop()
}

func createAuthConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, err
}

func createServerConn(interceptor *authclient.AuthInterceptor) (*grpc.ClientConn, error) {
	serverConn, err := grpc.Dial("localhost:8080",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		return nil, err
	}

	return serverConn, err
}
