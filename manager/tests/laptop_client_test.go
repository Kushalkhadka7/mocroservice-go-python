package laptopservice_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	database "manager/database"
	laptop "manager/pb"
	"manager/server"
	laptopservice "manager/service/laptop"
	laptopstore "manager/store"
	"testing"
)

var dbPort = 27017
var dbURI = "mongodb://localhost:27017"
var dbName = "mircoservices"
var dbPassword = "Kushal123"

func TestClientCreateLaptop(t *testing.T) {
	_, serverAddr := startTestLaptopServer(t)
	laptopClient := startGrpcClient(t, serverAddr)

	req := &laptop.Hello{Hello: "kushal"}
	callClient, err := laptopClient.SayHello(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, callClient)

	fmt.Println(callClient)

}

func startGrpcClient(t *testing.T, serverAddr string) laptop.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	require.NoError(t, err)

	return laptop.NewLaptopServiceClient(conn)
}

func startTestLaptopServer(t *testing.T) (*server.LaptopServer, string) {
	db, err := initializeDatabase(t)
	if err != nil {
		panic(err)
	}

	laptopStore := laptopstore.NewLaptopStore(db)
	laptopService := laptopservice.New(laptopStore)
	laptopServer := server.NewLaptopServer(laptopService, nil)

	server := server.New(8081)
	listener, err := server.StartHTTPServer()

	// Initializes new grpc server.
	grpcServer := grpc.NewServer()
	laptop.RegisterLaptopServiceServer(grpcServer, laptopServer)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func initializeDatabase(t *testing.T) (*mongo.Database, error) {
	monogDB := database.New(dbPort, dbURI, dbName, dbPassword)
	dbClient, err := monogDB.CreateClient()
	if err != nil {
		return nil, err
	}

	db := monogDB.InitializeDatabase(dbClient)

	return db, nil
}
