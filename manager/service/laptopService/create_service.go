package laptopservice

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	laptop "manager/pb"
	"manager/sample"
	storageservice "manager/service/storageService"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const maxImageSize = 1 << 20

// LaptopServer struct.
type LaptopServer struct {
	LaptopStore LaptopStore
	ImageStore  storageservice.Store
}

// NewLaptopService create new laptop store.
func NewLaptopService(laptopStore LaptopStore, imageStore storageservice.Store) *LaptopServer {
	return &LaptopServer{LaptopStore: laptopStore, ImageStore: imageStore}
}

// CreateLaptop creates new laptop.
func (service *LaptopServer) CreateLaptop(ctx context.Context, req *laptop.CreateLaptopRequest) (*laptop.CreateLaptopResponse, error) {
	data := req.GetLaptop()
	_, err := service.LaptopStore.Save(ctx, data)
	if err != nil {
		panic(err)
	}

	return &laptop.CreateLaptopResponse{}, nil
}

// SayHello hello func.
func (service *LaptopServer) SayHello(ctx context.Context, req *laptop.Hello) (*laptop.CreateLaptopResponse, error) {
	sample := sample.NewLaptop()

	result, err := service.LaptopStore.Save(context.Background(), sample)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sample.Id = result.InsertedID.(primitive.ObjectID).Hex()
	log.Println(sample)

	return &laptop.CreateLaptopResponse{Laptop: sample}, nil
}

// FetchAllLaptops fetch all laptops form the data base.
func (service *LaptopServer) FetchAllLaptops(void *laptop.Void, stream laptop.LaptopService_FetchAllLaptopsServer) error {
	log.Println("hello i am called.")
	result, err := service.LaptopStore.FetchAll(stream.Context())

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Print(result)

	if err := stream.Send(result); err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Successfully send laptop with id: %s", result.Laptop.GetId())

	return nil
}

// UploadLaptopImage upload image to server.
func (service *LaptopServer) UploadLaptopImage(stream laptop.LaptopService_UploadLaptopImageServer) error {
	fmt.Println("hello i am called here")

	req, err := stream.Recv()
	if err != nil {
		fmt.Println(err)
		return err
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("Request received with image type : %s and laptop id :%s", imageType, laptopID)

	result, err := service.LaptopStore.FindLaptop(context.Background(), laptopID)
	if err != nil {
		return err
	}
	if result == "" {
		return fmt.Errorf("No data found with id %s", laptopID)
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		fmt.Println("Starting to receive chunks of image data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("No more image data")
			break
		}
		if err != nil {
			return err
		}

		chunk := req.GetChunkData()
		imageSize += len(chunk)
		if imageSize > maxImageSize {
			return fmt.Errorf("Image size exceeds. %d", maxImageSize)
		}

		_, err = imageData.Write(chunk)
		if err != nil {
			return err
		}
	}

	imageID, err := service.ImageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		return err
	}
	if imageID == "" {
		return fmt.Errorf("No image is saved %s", "err")
	}

	err = service.LaptopStore.UpdateLaptop(laptopID, imageID)
	if err != nil {
		return err
	}

	res := &laptop.UploadImageResponse{
		ImageId:   imageID,
		LaptopId:  laptopID,
		ImageSize: uint32(imageSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return err
	}

	return nil
}
