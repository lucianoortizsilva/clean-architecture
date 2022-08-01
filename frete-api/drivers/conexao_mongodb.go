package drivers

import (
	"context"
	"log"

	"github.com/lucianoortizsilva/frete/frete-api/drivers/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConectarServidorMongoDB() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(configuration.URL_CONEXAO_MONGODB)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}
