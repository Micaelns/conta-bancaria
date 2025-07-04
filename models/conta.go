package models

import (
	"fmt"
	"time"
)

type Conta struct {
	ID       string  `bson:"_id,omitempty"`
	Cliente  string  `bson:"cliente"`
	Numero   string  `bson:"numero"`
	Agencia  string  `bson:"agencia"`
	ChavePix string  `bson:"chavePix"`
	Saldo    float64 `bson:"saldo"`
}

func (c *Conta) Depositar(valor float64) (*Operacao, error) {
	if valor <= 0 {
		return nil, fmt.Errorf(" R$ %.2f (valor inválido)",valor)
	}
	c.Saldo += valor
	return &Operacao{
		NumeroConta: c.Numero,
		Tipo:        string(Deposito),
		Valor:       valor,
		Timestamp:   time.Now(),
	},nil
}

func (c *Conta) FazerPix(valor float64) (*Operacao, error) {
	if valor <= 0 {
		return nil, fmt.Errorf(" R$ %.2f (valor inválido)",valor)
	}
	c.Saldo += valor
	return &Operacao{
		NumeroConta: c.Numero,
		Tipo:        string(Pix),
		Valor:       valor,
		Timestamp:   time.Now(),
	},nil
}

func (c *Conta) Sacar(valor float64) (*Operacao, error) {
	if valor > c.Saldo {
		err := fmt.Errorf(" R$ %.2f (saldo insuficiente)",c.Saldo)
		return nil,err
	}
	c.Saldo -= valor
	return &Operacao{
		NumeroConta: c.Numero,
		Tipo:        string(Saque),
		Valor:       valor,
		Timestamp:   time.Now(),
	},nil
}

