package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		uri:        "mongodb://localhost:27017/microservice",
		dbName:     "microservice",
		dbPassword: "Kushal123",
	}
}

// CreateClient initializes new mongo db client.
func (db *MongoDB) CreateClient() (*mongo.Client, error) {
	clientOptins := options.Client().ApplyURI(db.uri)

	client, err := mongo.Connect(context.TODO(), clientOptins)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// InitializeDatabase mongo database.
func (db *MongoDB) InitializeDatabase(dbClient *mongo.Client) *mongo.Database {
	databaseInstance := dbClient.Database(db.dbName)

	return databaseInstance
}

// CloseDBConnection mongo db connection.
func (db *MongoDB) CloseDBConnection(dbClient *mongo.Client) error {
	err := dbClient.Disconnect(context.TODO())

	if err != nil {
		return err
	}

	return nil
}
