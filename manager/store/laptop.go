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
	Save(ctx context.Context, laptop *laptop.Laptop) (string, error)
	FetchAll(ctx context.Context) (*laptop.FetchLaptopResposne, error)
	FindLaptop(ctx context.Context, laptopID string) (string, error)
	UpdateLaptop(ctx context.Context, laptopID string, imageID string) error
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
func (store *CreateLaptopStore) FetchAll(ctx context.Context) (*laptop.FetchLaptopResposne, error) {
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

// FindLaptop find laptop with given id
func (store *CreateLaptopStore) FindLaptop(ctx context.Context, laptopID string) (string, error) {
	ID := "5f57b3559c4bf7fc1dd42299"

	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return "", err
	}

	collection := store.Collection("laptop")

	result := collection.FindOne(ctx, bson.M{"_id": oid})
	if result == nil {
		return "", nil
	}

	var data primitive.M
	if err := result.Decode(&data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return ID, nil
}

// UpdateLaptop updates the laptop image.
func (store *CreateLaptopStore) UpdateLaptop(ctx context.Context, laptopID string, imageID string) error {
	collection := store.Collection("laptop")

	oid, err := primitive.ObjectIDFromHex(laptopID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": bson.M{
			"$eq": oid, // check if bool field has value of 'false'
		},
	}
	updateData := bson.M{"$set": bson.M{"imageId": imageID}}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
