package main

import (
	"conta-bancaria/conta"
	"fmt"
)

func main() {
	// Criar conta
	c := conta.Conta{
		Cliente:  "João Silva",
		Numero:   "12345-6",
		Agencia:  "0001",
		ChavePix: "joao@email.com",
	}

	// Dois depósitos de R$400
	c.Depositar(400)
	c.Depositar(400)

	// Mostrar extrato
	c.Extrato()

	// Sacar R$300
	if err := c.Sacar(300); err != nil {
		fmt.Println("Erro no saque:", err)
	}

	// Tentar sacar R$900 (deve falhar)
	if err := c.Sacar(900); err != nil {
		fmt.Println("Erro no saque:", err)
	}
}