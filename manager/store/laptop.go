package laptopstore

import (
	"context"
	"fmt"
	"log"
	laptop "manager/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// LaptopStore new.
type LaptopStore interface {
	FetchAll(ctx context.Context) ([]*laptop.Laptop, error)
	Save(ctx context.Context, laptop *laptop.Laptop) (string, error)
	FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error)
	UpdateLaptop(ctx context.Context, laptopID string, imageID string) (*mongo.UpdateResult,error)
}

// CreateLaptopStore new.
type CreateLaptopStore struct {
	*mongo.Database
}

// NewLaptopStore laptop store.
func NewLaptopStore(db *mongo.Database) LaptopStore {
	return &CreateLaptopStore{db}
}

// Save new laptop.
func (store *CreateLaptopStore) Save(ctx context.Context, req *laptop.Laptop) (string, error) {
	newData, err := bson.Marshal(req)

	collection := store.Collection("laptop")

	res, err := collection.InsertOne(context.Background(), newData)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// FetchAll fetch all laptops from db.
func (store *CreateLaptopStore) FetchAll(ctx context.Context) ([]*laptop.Laptop, error) {
	collection := store.Collection("laptop")
	var result []*laptop.Laptop

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {

	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var resposne *laptop.Laptop
		if err = cursor.Decode(&resposne); err != nil {
			log.Fatal(err)
		}

		fmt.Println(resposne)

		result = append(result, resposne)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// FindLaptop find laptop with given id
func (store *CreateLaptopStore) FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error) {
	collection := store.Collection("laptop")

	oid, err := primitive.ObjectIDFromHex(laptopID)
	if err != nil {
		return nil, err
	}

	fmt.Println("hello world")
	result := collection.FindOne(ctx, bson.M{"_id": oid})
	fmt.Println(result)
	if result == nil {
		return nil, fmt.Errorf("Unable to find laptop with id: %s", laptopID)
	}

	var response *laptop.Laptop
	if err := result.Decode(&response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

// UpdateLaptop updates the laptop image.
func (store *CreateLaptopStore) UpdateLaptop(ctx context.Context, laptopID string, imageID string) (*mongo.UpdateResult, error) {
	collection := store.Collection("laptop")

	oid, err := primitive.ObjectIDFromHex(laptopID)
	if err != nil {
		return nil,err
	}

	updateFilter := bson.M{
		"_id": bson.M{
			"$eq": oid, // check if bool field has value of 'false'
		},
	}
	toBeUpdatedPayload := bson.M{"$set": bson.M{"imageId": imageID}}

	result, err := collection.UpdateOne(ctx, updateFilter, toBeUpdatedPayload)
	if err != nil {
		return nil,err
	}

	return result,nil
}
