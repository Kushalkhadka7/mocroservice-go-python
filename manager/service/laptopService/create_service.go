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
func (service *LaptopServer) CreateLaptop(context.Context, *laptop.CreateLaptopRequest) (*laptop.CreateLaptopResponse, error) {
	fmt.Println("hell hello i am called and i am happy")
	return nil, fmt.Errorf("error from createlaptop service%s", "error")
}

// SayHello creates new laptop.
func (service *LaptopServer) SayHello(context.Context, *laptop.HelloWorld) (*laptop.HelloWorld, error) {
	data, err := service.Store.Save()
	if err != nil {
		return nil, err
	}

	return data, nil
}
