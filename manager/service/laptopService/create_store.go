package laptopservice

import (
	"context"
	"fmt"
	"log"
	laptop "manager/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// LaptopStore new.
type LaptopStore interface {
	Save(ctx context.Context, laptop *laptop.Laptop) (*mongo.InsertOneResult, error)
	FetchAll(ctx context.Context) (*laptop.FetchLaptopResposne, error)
}

// CreateLaptopStore new.
type CreateLaptopStore struct {
	*mongo.Database
}

// NewLaptopStore laptop store.
func NewLaptopStore(db *mongo.Database) *CreateLaptopStore {
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

// FetchAll fetch all laptops from db.
func (store *CreateLaptopStore) FetchAll(ctx context.Context) (*laptop.FetchLaptopResposne, error) {
	if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
		log.Printf("Error: %s", ctx.Err())
		return nil, ctx.Err()
	}

	collection := store.Collection("laptop")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result = &laptop.FetchLaptopResposne{}
		if err = cursor.Decode(&result.Laptop); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Laptop)

		return result, nil
	}

	return nil, nil
}
