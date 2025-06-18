package main

import (
	"conta-bancaria/conta"
	"fmt"
)

func main() {
	// Criar estrutura para armazenar múltiplas contas (mapa com chave = número da conta)
	contas := make(map[string]*conta.Conta)

	// Criar duas contas
	contas["123"] = &conta.Conta{
		Cliente:  "João Silva",
		Numero:   "123",
		Agencia:  "0001",
		ChavePix: "joao@email.com",
	}

	contas["456"] = &conta.Conta{
		Cliente:  "Maria Souza",
		Numero:   "456",
		Agencia:  "0002",
		ChavePix: "maria@celular.com",
	}

	// Operações na conta do João (123)
	contas["123"].ExtratoSimples()
	contas["123"].Depositar(400)
	contas["123"].Depositar(400)
	contas["123"].Sacar(300)
	if err := contas["123"].Sacar(900); err != nil {
		fmt.Println("Erro no saque:", err)
	}
	contas["123"].ExtratoCompleto()

	// Operações na conta da Maria (456)
	contas["456"].ExtratoSimples()
	contas["456"].Depositar(1000)
	contas["456"].Sacar(200)
	contas["456"].ExtratoCompleto()
}