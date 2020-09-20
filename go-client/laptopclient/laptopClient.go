package laptopclient

import (
	"bufio"
	pb "client/pb"
	"client/sample"
	"context"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

// LaptopClient initializes new laptop client.
type LaptopClient struct {
	context context.Context
	service pb.LaptopServiceClient
}

// Creator interface for laptop client.
type Creator interface {
	FetchAllLaptop()
	CreateNewLaptop() *pb.Laptop
	UploadLaptopImage(createdLaptop *pb.Laptop)
}

// NewLapotpClient creates new laptop client.
func NewLapotpClient(conn *grpc.ClientConn) Creator {
	service := pb.NewLaptopServiceClient(conn)

	return &LaptopClient{
		service: service,
		context: context.Background(),
	}
}

// CreateNewLaptop creates new laptop.
func (client *LaptopClient) CreateNewLaptop() *pb.Laptop {
	req := &pb.CreateLaptopRequest{
		Laptop: sample.NewLaptop(),
	}
	res, err := client.service.CreateLaptop(client.context, req)
	if err != nil {
		panic(err)
	}

	log.Println(res.Laptop)
	log.Println(res.Laptop.GetXId())
	log.Println("Laptop created successfully")

	return res.Laptop
}

// FetchAllLaptop retrives all laptops from db.
func (client *LaptopClient) FetchAllLaptop() {
	req := &pb.Void{}

	stream, err := client.service.FetchAllLaptops(client.context, req)
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

// UploadLaptopImage uploades new laptop image.
func (client *LaptopClient) UploadLaptopImage(createdLaptop *pb.Laptop) {
	imagePath := "./temp/laptop.jpg"

	imageFile, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("Cannot open image file :", err)
	}
	defer imageFile.Close()

	stream, err := client.service.UploadLaptopImage(context.Background())
	if err != nil {
		log.Fatal("Unable to establish server connection", err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
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

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
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
