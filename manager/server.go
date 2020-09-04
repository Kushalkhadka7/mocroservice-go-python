package main

import (
	"fmt"
	"manager/database"
	laptop "manager/pb"
	laptopservice "manager/service/laptopService"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server struct.
type Server struct {
	Port int
}

// NewServer creates a new server.
func NewServer(port int) *Server {
	return &Server{Port: port}
}

// StartServer starts the http server on given port.
func (server *Server) StartServer() (net.Listener, error) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		return nil, fmt.Errorf("Cannot establish http server connection:%v", err)
	}

	fmt.Printf("Server listening on port:%d", server.Port)

	return listener, nil
}

// StartGrpcServer starts a new grpc server.
func (server *Server) StartGrpcServer(listener net.Listener) error {

	monogDB := database.New()
	dbClient, err := monogDB.CreateClient()
	if err != nil {
		return err
	}

	db := monogDB.InitializeDatabase(dbClient)


	laptopStore := laptopservice.NewLaptopStore(db)
	laptopService := laptopservice.NewLaptopService(laptopStore)

	// Initializes new grpc server.
	grpcServer := grpc.NewServer()

	// Register laptop server to grpc server.
	laptop.RegisterLaptopServiceServer(grpcServer, laptopService)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
