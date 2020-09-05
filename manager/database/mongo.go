package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB struct the database.
type MongoDB struct {
	*mongo.Database
	port       int
	uri        string
	dbName     string
	dbPassword string
}

// New initializes new mongo db instance.
func New() *MongoDB {
	return &MongoDB{
		port:       27017,
		uri:        "mongodb://localhost:27017",
		dbName:     "mircoservices",
		dbPassword: "Kushal123",
	}
}

// CreateClient initializes new mongo db client.
func (db *MongoDB) CreateClient() (*mongo.Client, error) {
	clientOptins := options.Client().ApplyURI(db.uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptins)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("\nConnected to MongoDB on port : 27017")
	return client, nil
}

// InitializeDatabase mongo database.
func (db *MongoDB) InitializeDatabase(dbClient *mongo.Client) *mongo.Database {
	return dbClient.Database("microservices")
}

// CloseDBConnection mongo db connection.
func (db *MongoDB) CloseDBConnection(dbClient *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := dbClient.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
