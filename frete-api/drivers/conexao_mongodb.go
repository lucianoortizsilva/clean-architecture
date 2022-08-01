package drivers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConectarServidorMongoDB() (*mongo.Client, error) {

	var URL_CONEXAO_MONGODB = "mongodb://localhost:27017/frete-db"

	clientOptions := options.Client().ApplyURI(URL_CONEXAO_MONGODB)
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
