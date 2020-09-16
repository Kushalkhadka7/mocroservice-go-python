package laptopservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	laptop "manager/pb"
	"manager/store"
)

// LaptopService interface to implement laptop service.
type LaptopService interface {
	SaveLaptop(ctx context.Context, laptop *laptop.Laptop) (string, error)
	FetchAll(ctx context.Context) (*laptop.FetchLaptopResposne, error)
}

// CreateLaptopService creates new laptop service.
type CreateLaptopService struct {
	LaptopStore laptopstore.LaptopStore
}

// New cretes new service with latpop store.
func New(store laptopstore.LaptopStore) LaptopService {
	return &CreateLaptopService{LaptopStore: store}
}

// SaveLaptop service.
func (service *CreateLaptopService) SaveLaptop(ctx context.Context, laptop *laptop.Laptop) (string, error) {

	result, err := service.LaptopStore.Save(ctx, laptop)
	if err != nil {
		return "", status.Errorf(codes.Internal, "Cannot save laptop to db. %v", err)
	}

	return result,nil
}

// FetchAll laptops in db.
func (service *CreateLaptopService) FetchAll(ctx context.Context) (*laptop.FetchLaptopResposne, error) {
	if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
		log.Printf("Error: %s", ctx.Err())
		return nil, ctx.Err()
	}

	res, err := service.LaptopStore.FetchAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to fetch laptps %v", err)
	}

	return res, nil
}
