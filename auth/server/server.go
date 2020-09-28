package server

import (
	"auth/service"
	"auth/store"
	"fmt"
	"time"

	"net"

	pb "auth/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server takes port and initialize http server on the port.
type Server struct {
	Port int
}

// Creator interface creates and initializes http and grpc server.
type Creator interface {
	StartHTTPServer() (net.Listener, error)
	StartGrpcServer(listener net.Listener) error
}

// New creates a new server with given port.
func New(port int) Creator {
	return &Server{Port: port}
}

// StartHTTPServer starts the http server on given port.
func (server *Server) StartHTTPServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		return nil, fmt.Errorf("Cannot establish http server connection:%v", err)
	}

	fmt.Printf("Server listening on port:%d", server.Port)

	return listener, nil
}

// StartGrpcServer starts a new grpc server.
func (server *Server) StartGrpcServer(listener net.Listener) error {

	jwtManager := service.NewJWTMnanager("Kushal", 30*time.Second)

	userStore := store.NewUesrStore()
	userService := service.NewUserService(userStore, jwtManager)

	authServer := NewAuthServer(userService)

	// Initializes new grpc server.
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
