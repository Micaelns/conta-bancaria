package conta

import (
	"errors"
	"fmt"
)

// Conta representa uma conta bancária.
type Conta struct {
	Cliente  string
	Numero   string
	Agencia  string
	ChavePix string
	Saldo    float64
}

// Depositar adiciona valor ao saldo da conta.
func (c *Conta) Depositar(valor float64) {
	c.Saldo += valor
	fmt.Printf("Depósito de R$%.2f realizado. Saldo atual: R$%.2f\n", valor, c.Saldo)
}

// Sacar tenta subtrair valor do saldo.
func (c *Conta) Sacar(valor float64) error {
	if valor > c.Saldo {
		return errors.New("saldo insuficiente")
	}
	c.Saldo -= valor
	fmt.Printf("Saque de R$%.2f realizado. Saldo atual: R$%.2f\n", valor, c.Saldo)
	return nil
}

// Extrato imprime os dados da conta e saldo.
func (c *Conta) Extrato() {
	fmt.Printf("\n--- Extrato ---\n")
	fmt.Printf("Cliente: %s\n", c.Cliente)
	fmt.Printf("Agência: %s | Conta: %s\n", c.Agencia, c.Numero)
	fmt.Printf("Chave Pix: %s\n", c.ChavePix)
	fmt.Println()
	fmt.Printf("Saldo atual: R$%.2f\n", c.Saldo)
	fmt.Println("---------------")
	fmt.Println()
}