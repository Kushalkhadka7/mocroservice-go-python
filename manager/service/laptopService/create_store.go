package laptopservice

import (
	"context"
	laptop "manager/pb"

	"go.mongodb.org/mongo-driver/mongo"
)

// LaptopStore new.
type LaptopStore interface {
	Save(ctx context.Context, laptop *laptop.Laptop) (*laptop.Laptop, error)
}

// CreateLaptopStore new.
type CreateLaptopStore struct {
	*mongo.Database
	DB DBInterface
}

// DBInterface interface.
type DBInterface interface{}

// NewLaptopStore laptop store.
func NewLaptopStore(db DBInterface) LaptopStore {
	return &CreateLaptopStore{DB: db}
}

// Save new laptop.
func (store *CreateLaptopStore) Save(ctx context.Context, data *laptop.Laptop) (*laptop.Laptop, error) {
	return data, nil
}
