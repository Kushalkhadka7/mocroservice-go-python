package server

import (
	"auth/service"
	"auth/store"
	"fmt"

	"net"

	auth "auth/pbauth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const dbPort = 27017
const dbName = "mircoservices"
const dbPassword = "Kushal123"
const dbURI = "mongodb://localhost:27017"

// Server takes port and initialize http server on the port.
type Server struct {
	Port int
}

// Creator interface creates and initializes http and grpc server.
type Creator interface {
	StartHTTPServer() (net.Listener, error)
	// DBConnection() (*mongo.Database, error)
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
	// db, err := server.DBConnection()
	// if err != nil {
	// 	return err
	// }

	userStore := store.NewUesrStore()
	userService := service.NewUserService(userStore)

	authServer := NewAuthServer(userService)

	// // Initializes new grpc server.
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

// DBConnection initializes new database connection.
// func (server *Server) DBConnection() (*mongo.Database, error) {
// Creates database instance on the give port.
// monogDB := database.New(dbPort, dbURI, dbName, dbPassword)
// dbClient, err := monogDB.CreateClient()
// if err != nil {
// return nil, err
// }
//
// db := monogDB.InitializeDatabase(dbClient)
//
// return db, nil
// }
