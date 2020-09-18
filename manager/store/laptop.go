package laptopstore

import (
	"context"
	"google.golang.org/grpc/codes"
	laptop "manager/pb"
	"manager/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// LaptopStore new.
type LaptopStore interface {
	FetchAll(ctx context.Context) ([]*laptop.Laptop, error)
	Save(ctx context.Context, laptop *laptop.Laptop) (string, error)
	FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error)
	UpdateLaptop(ctx context.Context, laptopID string, imageID string) (*mongo.UpdateResult, error)
}

// CreateLaptopStore new.
type CreateLaptopStore struct {
	*mongo.Database
}

// NewLaptopStore laptop store.
func NewLaptopStore(db *mongo.Database) LaptopStore {
	return &CreateLaptopStore{db}
}

// Save new laptop to databse.
func (store *CreateLaptopStore) Save(ctx context.Context, req *laptop.Laptop) (string, error) {
	newData, err := bson.Marshal(req)
	collection := store.Collection("laptop")

	res, err := collection.InsertOne(context.Background(), newData)
	if err != nil {
		return "", util.Error(codes.Internal, "Unable to save laptop to db", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// FetchAll fetches all laptops available in database.
func (store *CreateLaptopStore) FetchAll(ctx context.Context) ([]*laptop.Laptop, error) {
	var result []*laptop.Laptop
	collection := store.Collection("laptop")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, util.Error(codes.Internal, "Unable to fetch data from db", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var resposne *laptop.Laptop
		if err = cursor.Decode(&resposne); err != nil {
			return nil, util.Error(codes.Internal, "Unable to decode data from db", err)
		}

		result = append(result, resposne)
	}
	if err := cursor.Err(); err != nil {
		return nil, util.Error(codes.Internal, "Cursor error", err)
	}

	return result, nil
}

// FindLaptop find laptop with given id
func (store *CreateLaptopStore) FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error) {
	collection := store.Collection("laptop")

	oid, err := util.GenerateOID(laptopID)
	if err != nil {
		return nil, err
	}

	result := collection.FindOne(ctx, bson.M{"_id": oid})
	if result == nil {
		return nil, util.Error(codes.NotFound, "Unable to find laptop", err)
	}

	var response *laptop.Laptop
	if err := result.Decode(&response); err != nil {
		return nil, util.Error(codes.Internal, "Unable to decode data from db", err)
	}

	return response, nil
}

// UpdateLaptop updates the laptop image.
func (store *CreateLaptopStore) UpdateLaptop(ctx context.Context, laptopID string, imageID string) (*mongo.UpdateResult, error) {
	collection := store.Collection("laptop")

	oid, err := util.GenerateOID(laptopID)
	if err != nil {
		return nil, err
	}

	updateFilter := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}
	toBeUpdatedPayload := bson.M{"$set": bson.M{"imageId": imageID}}

	result, err := collection.UpdateOne(ctx, updateFilter, toBeUpdatedPayload)
	if err != nil {
		return nil, util.Error(codes.Internal, "Unable to update data", err)
	}

	return result, nil
}
