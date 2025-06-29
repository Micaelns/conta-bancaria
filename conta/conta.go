package conta

import (
	"fmt"
	"time"
)

type Conta struct {
	cliente  string
	numero   string
	agencia  string
	chavePix string 
	saldo 	 float64 // agora é privado (visível apenas no pacote `conta`)
	operacoes  []Operacao
}

func NovaConta(cliente, numero, agencia, chavePix string) *Conta {
	return &Conta{
		cliente:  cliente,
		numero:   numero,
		agencia:  agencia,
		chavePix: chavePix,
	}
}

func (c *Conta) GetCliente() string  { return c.cliente }
func (c *Conta) GetNumero() string   { return c.numero }
func (c *Conta) GetAgencia() string  { return c.agencia }
func (c *Conta) GetChavePix() string { return c.chavePix }
func (c *Conta) GetSaldo() float64   { return c.saldo }


func (c *Conta) Depositar(valor float64) {
	c.saldo += valor
	c.registrarOperacao(Deposito, valor, nil)
}

func (c *Conta) FazerPix(valor float64) {
	c.saldo += valor
	c.registrarOperacao(Pix, valor, nil)
}

func (c *Conta) Sacar(valor float64) error {
	if valor > c.saldo {
		err := fmt.Errorf("R$%.2f (saldo insuficiente)",c.saldo)
		c.registrarOperacao(Saque, valor, err)
		return err
	}
	c.saldo -= valor
	c.registrarOperacao(Saque, valor, nil)
	return nil
}

// ---------- MÉTODO INTERNO (privado) ----------
//nome do método iniciado por letra minúscula método privado

func (c *Conta) registrarOperacao(tipo TipoOperacao, valor float64, err error) {
	op := Operacao{
		Tipo:      tipo,
		Valor:     valor,
		DataHora: time.Now(),
	}
	if err != nil {
		op.Erro = err.Error()
	}
	c.operacoes = append(c.operacoes, op)
}