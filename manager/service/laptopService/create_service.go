package laptopservice

import (
	"context"
	"fmt"

	laptop "manager/pb"
)

// LaptopServer struct.
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopService create new laptop store.
func NewLaptopService(store LaptopStore) laptop.LaptopServiceServer {
	fmt.Println("hello i am connected to grpc server")
	return &LaptopServer{Store: store}
}

// CreateLaptop creates new laptop.
func (service *LaptopServer) CreateLaptop(ctx context.Context, req *laptop.CreateLaptopRequest) (*laptop.CreateLaptopResponse, error) {
	data := req.GetLaptop()
	result, err := service.Store.Save(ctx, data)
	if err != nil {
		panic(err)
	}

	return &laptop.CreateLaptopResponse{Laptop: result}, nil
}
