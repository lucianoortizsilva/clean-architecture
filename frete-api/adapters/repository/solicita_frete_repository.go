package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SolicitaFreteDB struct {
	client *mongo.Client
}

func NewSolicitaFreteDB(clienteMongoDB *mongo.Client) *SolicitaFreteDB {
	return &SolicitaFreteDB{client: clienteMongoDB}
}

func (sf *SolicitaFreteDB) Insert(PedidoId string, Regiao string, PesoTotalPedido float64) error {

	fretes := sf.client.Database("frete-db").Collection("fretes")

	document := bson.D{
		{Key: "PedidoId", Value: PedidoId},
		{Key: "Regiao", Value: Regiao},
		{Key: "PesoTotalPedido", Value: PesoTotalPedido},
		{Key: "CriadoEm", Value: time.Now()},
	}

	context, _ := context.WithTimeout(context.Background(), 15*time.Second)

	doc, err := fretes.InsertOne(context, document)

	if err != nil {
		fmt.Println(doc)
	}

	return err
}
