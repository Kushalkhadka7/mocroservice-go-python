package laptopservice

import (
	"fmt"
	laptop "manager/pb"

	"go.mongodb.org/mongo-driver/mongo"
)

// LaptopStore new.
type LaptopStore interface {
	Save() (*laptop.HelloWorld, error)
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
func (store *CreateLaptopStore) Save() (*laptop.HelloWorld, error) {
	fmt.Println("hell hello i am called and i am happy")

	return &laptop.HelloWorld{Hello: "hello howr are you"}, nil
}
