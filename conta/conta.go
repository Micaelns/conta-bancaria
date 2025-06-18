package conta

import (
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
	fmt.Printf("[Conta %s] + R$%.2f\n", c.Numero, valor)
}

// Sacar tenta subtrair valor do saldo.
func (c *Conta) Sacar(valor float64) error {
	if valor > c.Saldo {
		return fmt.Errorf("[Conta %s] - R$%.2f (saldo insuficiente)",c.Numero, valor)
	}
	c.Saldo -= valor
	fmt.Printf("[Conta %s] - R$%.2f\n", c.Numero,  valor)
	return nil
}

// Imprime saldo.
func (c *Conta) ExtratoSimples() {
	fmt.Printf("\n--- Saldo ---\n")
	fmt.Printf("Agência: %s | Conta: %s\n", c.Agencia, c.Numero)
	fmt.Printf("Saldo atual: R$%.2f\n", c.Saldo)
	fmt.Println("---------------")
	fmt.Println()
}

// Extrato imprime os dados da conta e saldo.
func (c *Conta) ExtratoCompleto() {
	fmt.Printf("\n--- Extrato ---\n")
	fmt.Printf("Cliente: %s\n", c.Cliente)
	fmt.Printf("Agência: %s | Conta: %s\n", c.Agencia, c.Numero)
	fmt.Printf("Chave Pix: %s\n", c.ChavePix)
	fmt.Println()
	fmt.Printf("Saldo atual: R$%.2f\n", c.Saldo)
	fmt.Println("---------------")
	fmt.Println()
}