package main

import (
	"bufio"
	authclient "client/authClient"
	laptop "client/pb"
	auth "client/pbauth"
	"client/sample"
	"context"
	"fmt"
	"io"
	"log"
	"os"
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

	laptopClient := laptop.NewLaptopServiceClient(serverConn)

	createNewLaptop(context.Background(), laptopClient)
}

func createUser(ctx context.Context, authClient auth.AuthServiceClient) {
	req := &auth.CreateUserRequest{
		User: &auth.User{
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

// createNewLaptop creates new laptop.
func createNewLaptop(ctx context.Context, client laptop.LaptopServiceClient) *laptop.Laptop {
	req := &laptop.CreateLaptopRequest{
		Laptop: sample.NewLaptop(),
	}
	res, err := client.CreateLaptop(context.Background(), req)
	if err != nil {
		panic(err)
	}

	log.Println(res.Laptop)
	log.Println(res.Laptop.GetXId())
	log.Println("Laptop created successfully")

	return res.Laptop
}

// fetchAllLaptop retrives all laptops from db.
func fetchAllLaptop(ctx context.Context, laptopClient laptop.LaptopServiceClient) {
	req := &laptop.Void{}

	stream, err := laptopClient.FetchAllLaptops(context.Background(), req)
	if err != nil {
		panic(err)
	}

	res, err := stream.Recv()
	if err != nil {
		panic(err)
	}

	log.Println(res.Laptop)
	log.Println("Laptop created successfully")
}

func uploadLaptopImage(ctx context.Context, laptopClient laptop.LaptopServiceClient, createdLaptop *laptop.Laptop) {
	imagePath := "./temp/laptop.jpg"

	imageFile, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("Cannot open image file :", err)
	}
	defer imageFile.Close()

	stream, err := laptopClient.UploadLaptopImage(ctx)
	if err != nil {
		log.Fatal("Unable to establish server connection", err)
	}

	req := &laptop.UploadImageRequest{
		Data: &laptop.UploadImageRequest_Info{
			Info: &laptop.ImageInfo{
				LaptopId:  createdLaptop.GetXId(),
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

	log.Println(res)
	log.Println("Sucessfully uploaded image")
}
