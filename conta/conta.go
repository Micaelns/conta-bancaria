package conta

import (
	"fmt"
	"time"
)

type TipoOperacao string

const (
	Deposito TipoOperacao = "Depósito"
	Saque    TipoOperacao = "Saque"
)

type Operacao struct {
	Tipo      TipoOperacao
	Valor     float64
	Timestamp time.Time
	Erro      string // se falhar, registra o motivo
}

// Conta representa uma conta bancária.
type Conta struct {
	Cliente  string
	Numero   string
	Agencia  string
	ChavePix string
	Saldo    float64
	Operacoes  []Operacao
}

func (c *Conta) registrarOperacao(tipo TipoOperacao, valor float64, err error) {
	op := Operacao{
		Tipo:      tipo,
		Valor:     valor,
		Timestamp: time.Now(),
	}
	if err != nil {
		op.Erro = err.Error()
	}
	c.Operacoes = append(c.Operacoes, op)
}

// Depositar adiciona valor ao saldo da conta.
func (c *Conta) Depositar(valor float64) {
	c.Saldo += valor
	fmt.Printf("[Conta %s] + R$%.2f\n", c.Numero, valor)
	c.registrarOperacao(Deposito, valor, nil)
}

// Sacar tenta subtrair valor do saldo.
func (c *Conta) Sacar(valor float64) error {
	if valor > c.Saldo {
		err := fmt.Errorf("- R$%.2f (saldo insuficiente)",valor)
		c.registrarOperacao(Saque, valor, err)
		return err
	}
	c.Saldo -= valor
	fmt.Printf("[Conta %s] - R$%.2f\n", c.Numero,  valor)
	c.registrarOperacao(Saque, valor, nil)
	return nil
}

// Imprime saldo.
func (c *Conta) ExtratoSimples() {
	fmt.Printf("\n--- Saldo ---\n")
	fmt.Printf("Cliente: %s\n", c.Cliente)
	fmt.Printf("Agência: %s | Conta: %s\n", c.Agencia, c.Numero)
	fmt.Printf("Saldo atual: R$%.2f\n", c.Saldo)
	fmt.Println("---------------")
	fmt.Println()
}

// Extrato imprime os dados da conta e saldo.
func (c *Conta) ExtratoCompleto() {
	fmt.Println("\n--- Extrato ---")
	fmt.Printf("Cliente: %s\n", c.Cliente)
	fmt.Printf("Agência: %s | Conta: %s\n", c.Agencia, c.Numero)
	for _, op := range c.Operacoes {
		status := ""
		if op.Erro != "" {
			status = "— FALHA: " + op.Erro
		}
		sinal := "-"
		if op.Tipo == Deposito {
			sinal = "+"
		}
		fmt.Printf("[%s] %s de R$%.2f %s\n",
			op.Timestamp.Format("02/01/2006 15:04:05"),
			sinal,
			op.Valor,
			status,
		)
	}
	fmt.Printf("Saldo atual: R$%.2f\n", c.Saldo)
}