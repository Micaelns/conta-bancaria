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

func (c *Conta) Depositar(valor float64) *Operacao {
	op := &Operacao{
		NumeroConta: c.Numero,
		Tipo:        string(Deposito),
		Valor:       valor,
		Timestamp:   time.Now(),
	}
	if valor <= 0 {
		op.Erro = fmt.Sprintf(" R$ %.2f (valor inválido)",valor)
		return op
	}
	c.Saldo += valor
	return op
}

func (c *Conta) FazerPix(valor float64) *Operacao {
	op := &Operacao{
		NumeroConta: c.Numero,
		Tipo:        string(Pix),
		Valor:       valor,
		Timestamp:   time.Now(),
	}
	if valor <= 0 {
		op.Erro = fmt.Sprintf(" R$ %.2f (valor inválido)",valor)
		return op
	}
	c.Saldo += valor
	return op
}

func (c *Conta) Sacar(valor float64) *Operacao {
	op := &Operacao{
		NumeroConta: c.Numero,
		Tipo:        string(Saque),
		Valor:       valor,
		Timestamp:   time.Now(),
	}

	if valor < 0 || valor > c.Saldo {
		op.Erro = fmt.Sprintf(" R$ %.2f (valor inválido)",valor)
		return op
	}

	if valor < 0 || valor > c.Saldo {
		op.Erro = fmt.Sprintf(" R$ %.2f (saldo insuficiente)",c.Saldo)
		return op
	}
	c.Saldo -= valor
	return op
}

