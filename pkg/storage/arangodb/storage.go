package arangodb

import (
	"context"

	arangodbdriver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func NewArangoDBClient(
	host string,
	port string,
	user string,
	password string,
	database string,
) (*arangodbdriver.Database, error) {
	// Create a client
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + host + ":" + port},
	})
	if err != nil {
		return nil, err
	}
	client, err := arangodbdriver.NewClient(arangodbdriver.ClientConfig{
		Connection:     conn,
		Authentication: arangodbdriver.BasicAuthentication(user, password),
	})
	if err != nil {
		return nil, err
	}

	// Use the client to get a database
	db, err := client.Database(context.Background(), database)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
