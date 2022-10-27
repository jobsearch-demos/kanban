package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Storage storage implementation.
type Storage struct {
	Host     string `json:"host" default:"localhost" validate:"required, ipv4|ipv6|hostname"`
	Port     string `json:"port" default:"27017" validate:"required, gte=0, lte=65535"`
	User     string `json:"user" default:"root" validate:"required"`
	Password string `json:"password" default:"root" validate:"required"`
	Database string `json:"database" default:"search" validate:"required"`
}

func (s *Storage) GetDB() (*mongo.Database, error) {
	ctx := context.TODO()
	// prepare the connection string
	url := fmt.Sprint("mongodb://", s.User, ":", s.Password, "@", s.Host, ":", s.Port)

	// prepare the default options
	clientOptions := options.Client().ApplyURI(url)

	// create the client connection
	log.Printf("Connecting to mongo server ...\n")
	client, err := mongo.Connect(ctx, clientOptions)

	// if there was an error creating the client, return an error
	if err != nil {
		return nil, fmt.Errorf("could not get the mongo client %s: ", err)
	}

	// check if the client is alive
	log.Printf("Pinging the mongo server ...\n")
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("mongo heartbeat failed: %s", err.Error())
	}

	log.Printf("Connecting specified database ...\n")
	db := client.Database(s.Database)

	if db == nil {
		return nil, fmt.Errorf("could not connect to database named '%s': ", s.Database)
	}

	log.Printf("Connected to specified database \n")
	return db, nil
}
