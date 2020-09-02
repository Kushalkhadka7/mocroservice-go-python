package laptopservice

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// LaptopStore new.
type LaptopStore interface {
	Save() error
}

// LaptopInterface interface.
type LaptopInterface interface{}

// CreateLaptopStore new.
type CreateLaptopStore struct {
	*mongo.Database
	DB LaptopInterface
}

// NewLaptopStore laptop store.
func NewLaptopStore(db LaptopInterface) LaptopStore {
	return &CreateLaptopStore{DB: db}
}

// Save new laptop.
func (store *CreateLaptopStore) Save() error {
	result := struct {
		Name string
		id   string
	}{}

	data, err := store.Collection("laptop").Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	data.Decode(&result)

	fmt.Println("from save laptop")
	fmt.Println(result)

	return nil
}
