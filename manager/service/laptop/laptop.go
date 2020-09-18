package laptopservice

import (
	"context"
	"google.golang.org/grpc/codes"
	laptop "manager/pb"
	"manager/store"
	"manager/util"
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
		return "", err
	}

	return result, nil
}

// FetchAllLaptops laptops in db.
func (service *CreateLaptopService) FetchAllLaptops(ctx context.Context) ([]*laptop.Laptop, error) {
	// Checks context error.
	if ctx.Err() == context.Canceled {
		return nil, util.Error(codes.DeadlineExceeded, "Context cancelled", ctx.Err())
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, util.Error(codes.DeadlineExceeded, "Context deadline exceedes", ctx.Err())
	}

	result, err := service.LaptopStore.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindLaptop find laptop with given id
func (service *CreateLaptopService) FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error) {
	result, err := service.LaptopStore.FindLaptop(ctx, laptopID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateLaptop updates laptop in store.
func (service *CreateLaptopService) UpdateLaptop(ctx context.Context, laptopID string, imageID string) error {
	_, err := service.LaptopStore.UpdateLaptop(ctx, laptopID, imageID)
	if err != nil {
		return err
	}

	return nil
}
