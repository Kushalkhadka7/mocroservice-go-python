package server

import (
	"context"
	"fmt"
	"time"

	laptop "manager/pb"
	"manager/sample"
	imageservice "manager/service/image"
	laptopservice "manager/service/laptop"
)

// LaptopServer struct.
type LaptopServer struct {
	LaptopService laptopservice.LaptopService
	ImageService  imageservice.CreateImageStorage
}

// NewLaptopServer create new laptop store.
func NewLaptopServer(laptopStore laptopservice.LaptopService, imageStore imageservice.CreateImageStorage) *LaptopServer {
	return &LaptopServer{LaptopService: laptopStore, ImageService: imageStore}
}

// CreateLaptop creates new laptop.
func (service *LaptopServer) CreateLaptop(ctx context.Context, req *laptop.CreateLaptopRequest) (*laptop.CreateLaptopResponse, error) {
	sample := sample.NewLaptop()

	result, err := service.LaptopService.SaveLaptop(context.Background(), sample)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sample.XId = result

	return &laptop.CreateLaptopResponse{Laptop: sample}, nil
}

// SayHello hello func.
func (service *LaptopServer) SayHello(ctx context.Context, req *laptop.Hello) (*laptop.CreateLaptopResponse, error) {
		sample := sample.NewLaptop()

	result, err := service.LaptopService.SaveLaptop(context.Background(), sample)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sample.XId = result

	return &laptop.CreateLaptopResponse{Laptop: sample}, nil
}

// FetchAllLaptops fetch all laptops form the data base.
func (service *LaptopServer) FetchAllLaptops(void *laptop.Void, stream laptop.LaptopService_FetchAllLaptopsServer) error {

	result, err := service.LaptopService.FetchAllLaptops(stream.Context())
	if err != nil {
		return err
	}

	for i := range result {
		res := &laptop.FetchLaptopResposne{
			Laptop: result[i],
		}

		err := stream.Send(res)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("Successfully send laptop with id: %s", res.Laptop.GetBrand())
	}

	return nil
}

// UploadLaptopImage upload image to server.
func (service *LaptopServer) UploadLaptopImage(stream laptop.LaptopService_UploadLaptopImageServer) error {
	
	result, err := service.ImageService.Save(stream)
	if err !=nil {
		return err
	}
	
	err = stream.SendAndClose(result)
	if err != nil {
		return err
	}

	return nil
}
