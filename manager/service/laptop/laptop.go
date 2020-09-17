package laptopservice

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	laptop "manager/pb"
	"manager/store"
)

// LaptopService interface to implement laptop service.
type LaptopService interface {
	FetchAllLaptops(ctx context.Context) ([]*laptop.Laptop, error)
	SaveLaptop(ctx context.Context, laptop *laptop.Laptop) (string, error)
	FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error)
	UpdateLaptop(ctx context.Context, laptopID string, imageID string) error
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

// FetchAllLaptops laptops in db.
func (service *CreateLaptopService) FetchAllLaptops(ctx context.Context) ([]*laptop.Laptop, error) {
	// Checks context error.
	if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
		log.Printf("Error: %s", ctx.Err())
		return nil, ctx.Err()
	}

	result, err := service.LaptopStore.FetchAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to fetch laptps %v", err)
	}
	
	return result, nil
}

// FindLaptop find laptop with given id
func (service *CreateLaptopService) FindLaptop(ctx context.Context,laptopID string) (*laptop.Laptop, error) {
	ID := "5f57b3559c4bf7fc1dd42299"

	result,err:= service.LaptopStore.FindLaptop(ctx,ID)
	if err !=nil {
		return nil,err
	}

	return result,nil
}

// UpdateLaptop updates laptop in store.
func (service *CreateLaptopService) UpdateLaptop(ctx context.Context, laptopID string, imageID string) error{
	ID := "5f57b3559c4bf7fc1dd42299"

	result,err:= service.LaptopStore.UpdateLaptop(ctx,ID,ID)
	fmt.Println(result)
	if err !=nil {
		return err
	}

	return nil
}