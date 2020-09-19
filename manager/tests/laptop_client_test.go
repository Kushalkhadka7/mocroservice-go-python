package laptopservice_test

import (
	"bufio"
	"context"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"io"
	"log"
	database "manager/database"
	laptop "manager/pb"
	sample "manager/sample"
	"manager/server"
	laptopservice "manager/service/laptop"
	laptopstore "manager/store"
	"os"
	"testing"
)

var dbPort = 27017
var dbName = "mircoservices"
var dbPassword = "Kushal123"
var dbURI = "mongodb://localhost:27017"

func TestClientCreateLaptop(t *testing.T) {
	// _, serverAddr := startTestLaptopServer(t)
	laptopClient := startGrpcClient(t, ":8080")

	// Test createLaptop function.
	createdLaptop := createLaptop(t, laptopClient)

	// Test fetch all laptops function.
	fetchAllLaptop(t, laptopClient)

	testUploadImage(t, laptopClient, createdLaptop.Laptop)
}

func testUploadImage(t *testing.T, laptopClient laptop.LaptopServiceClient, newLaptop *laptop.Laptop) {
	imagePath := "../tmp/laptop.jpg"

	imageFile, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file :", err)
	}
	defer imageFile.Close()

	stream, err := laptopClient.UploadLaptopImage(context.Background())
	if err != nil {
		log.Fatal("Unable to establish server connection", err)
	}

	req := &laptop.UploadImageRequest{
		Data: &laptop.UploadImageRequest_Info{
			Info: &laptop.ImageInfo{
				LaptopId:  newLaptop.XId,
				ImageType: "jpg",
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("Unable to send image info", err)
	}

	reader := bufio.NewReader(imageFile)
	buffer := make([]byte, 1024)

	for {
		data, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Unable to read image data.", err)
		}

		req := &laptop.UploadImageRequest{
			Data: &laptop.UploadImageRequest_ChunkData{
				ChunkData: buffer[:data],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("unable to send image data.", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Unable to receive resposne", err)
	}

	log.Println("Sucessfully uploaded image")
	log.Println(res)
}

func fetchAllLaptop(t *testing.T, laptopClient laptop.LaptopServiceClient) {
	req := &laptop.Void{}

	res, err := laptopClient.FetchAllLaptops(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func createLaptop(t *testing.T, laptopClient laptop.LaptopServiceClient) *laptop.CreateLaptopResponse {
	sampleLaptop := sample.NewLaptop()
	req := &laptop.CreateLaptopRequest{Laptop: sampleLaptop}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)

	return res
}

// starts test grpc client for testing grpc server.
func startGrpcClient(t *testing.T, serverAddr string) laptop.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	require.NoError(t, err)

	return laptop.NewLaptopServiceClient(conn)
}

// startTestLaptopServer starts test server.
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

// Iniitlaizes test database connection.
func initializeDatabase(t *testing.T) (*mongo.Database, error) {
	monogDB := database.New(dbPort, dbURI, dbName, dbPassword)
	dbClient, err := monogDB.CreateClient()
	if err != nil {
		return nil, err
	}

	db := monogDB.InitializeDatabase(dbClient)

	return db, nil
}
