package repositories

import (
	"conta-bancaria/infra"
	"conta-bancaria/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OperacaoRepository struct {
	col *mongo.Collection
}

func NovaOperacaoRepository() *OperacaoRepository {
	return &OperacaoRepository{
		col: infra.OperacoesCollection(),
	}
}

func (r *OperacaoRepository) Registrar(op *models.Operacao) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.col.InsertOne(ctx, op)
	return err
}

func (r *OperacaoRepository) ListarPorConta(numeroConta string) ([]models.Operacao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"numeroConta": numeroConta}
	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var operacoes []models.Operacao
	for cursor.Next(ctx) {
		var op models.Operacao
		if err := cursor.Decode(&op); err != nil {
			return nil, err
		}
		operacoes = append(operacoes, op)
	}

	return operacoes, nil
}
