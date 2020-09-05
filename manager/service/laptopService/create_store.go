package laptopservice

import (
	"context"
	laptop "manager/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// LaptopStore new.
type LaptopStore interface {
	Save(ctx context.Context, laptop *laptop.Laptop) (*mongo.InsertOneResult, error)
}

// CreateLaptopStore new.
type CreateLaptopStore struct {
	*mongo.Database
}

// DBInterface interface.
type DBInterface interface{}

// NewLaptopStore laptop store.
func NewLaptopStore(db *mongo.Database) LaptopStore {
	return &CreateLaptopStore{db}
}

// Save new laptop.
func (store *CreateLaptopStore) Save(ctx context.Context, req *laptop.Laptop) (*mongo.InsertOneResult, error) {
	newData, err := bson.Marshal(req)

	collection := store.Collection("laptop")

	res, err := collection.InsertOne(context.Background(), newData)
	if err != nil {
		return nil, err
	}

	return res, nil
}
