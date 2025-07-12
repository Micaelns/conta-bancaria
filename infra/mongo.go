package infra

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConectarMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Testa a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	MongoClient = client
	fmt.Println("✅ Conectado ao MongoDB")
	return nil
}

func ContasCollection() *mongo.Collection {
	return MongoClient.Database("banco_contas").Collection("contas")
}

func OperacoesCollection() *mongo.Collection {
	return MongoClient.Database("banco_contas").Collection("operacoes")
}

func ObterColecao(nome string) *mongo.Collection {
	return MongoClient.Database("banco_contas").Collection(nome)
}
