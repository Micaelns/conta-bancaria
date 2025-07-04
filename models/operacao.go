package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TipoOperacao string

const (
	Deposito 	TipoOperacao = "Depósito"
	Pix    		TipoOperacao = "Pix"
	Saque    	TipoOperacao = "Saque"
)

type Operacao struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	NumeroConta string             `bson:"numeroConta"`
	Tipo        string             `bson:"tipo"`
	Valor       float64            `bson:"valor"`
	Timestamp   time.Time          `bson:"timestamp"`
	Erro        string             `bson:"erro,omitempty"`
}
