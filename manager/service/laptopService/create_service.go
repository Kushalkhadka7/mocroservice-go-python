package laptopservice

import (
	"context"
	"fmt"
	"log"

	laptop "manager/pb"
	"manager/sample"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LaptopServer struct.
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopService create new laptop store.
func NewLaptopService(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

// CreateLaptop creates new laptop.
func (service *LaptopServer) CreateLaptop(ctx context.Context, req *laptop.CreateLaptopRequest) (*laptop.CreateLaptopResponse, error) {
	data := req.GetLaptop()
	_, err := service.Store.Save(ctx, data)
	if err != nil {
		panic(err)
	}

	return &laptop.CreateLaptopResponse{}, nil
}

// SayHello hello func.
func (service *LaptopServer) SayHello(ctx context.Context, req *laptop.Hello) (*laptop.CreateLaptopResponse, error) {
	sample := sample.NewLaptop()

	result, err := service.Store.Save(context.Background(), sample)
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
	result, err := service.Store.FetchAll(stream.Context())

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
