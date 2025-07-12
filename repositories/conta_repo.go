package repositories

import (
	"conta-bancaria/infra"
	"conta-bancaria/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContaRepositoryInterface interface {
	BuscarConta(numero string) (*models.Conta, error)
	BuscarContaPix(chavePix string) (*models.Conta, error)
	Adicionar(conta *models.Conta) error
	AtualizarSaldo(numero string, novoSaldo float64) error
}

type ContaRepository struct {
	col *mongo.Collection
}

func NovaContaRepository() *ContaRepository {
	return &ContaRepository{
		col: infra.ContasCollection(),
	}
}

func (r *ContaRepository) Adicionar(c *models.Conta) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.col.InsertOne(ctx, c)
	return err
}

func (r *ContaRepository) BuscarConta(numero string) (*models.Conta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var conta models.Conta
	err := r.col.FindOne(ctx, bson.M{"numero": numero}).Decode(&conta)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("conta não encontrada")
		}
		return nil, err
	}

	return &conta, nil
}

func (r *ContaRepository) BuscarContaPix(chavePix string) (*models.Conta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var conta models.Conta
	err := r.col.FindOne(ctx, bson.M{"chavePix": chavePix}).Decode(&conta)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("conta não encontrada")
		}
		return nil, err
	}

	return &conta, nil
}

func (r *ContaRepository) AtualizarSaldo(numero string, novoSaldo float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"numero": numero}
	update := bson.M{"$set": bson.M{"saldo": novoSaldo}}

	result, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("conta não encontrada para atualizar saldo")
	}
	return nil
}
