package main

import (
	"conta-bancaria/conta"
	"encoding/json"
	"fmt"
)

func main() {
	repo := conta.NovoRepositorio()

	service := conta.NovoContaService(repo)

	// Criar contas
	service.CriarConta("João Silva", "123", "0001", "joao@email.com")
	service.CriarConta("Maria Souza", "456", "0002", "maria@celular.com")

	// Operações conta 123
	service.Depositar("123", 400)
	service.Depositar("123", 400)
	realizarSaque(service, "123", 300) 
	realizarSaque(service, "123", 900)
	mostrarJSON(service,"123") 

	// Operações conta 456
	service.Depositar("456", 1000)
	realizarSaque(service, "456", 200)
	mostrarJSON(service,"456") 

}

func realizarSaque(service *conta.ContaService, numConta string, valor float64) {
	if err := service.Sacar(numConta, valor); err != nil {
	 	fmt.Printf("Erro no saque [Conta: %s]: %v \n", numConta, err)
	} 
}

func mostrarJSON(service *conta.ContaService, numConta string) {
	if c, err := service.ConsultarConta(numConta); err == nil {
		fmt.Printf("[Conta: %s] \n",numConta)
		response := c.ToResponse()
		jsonData, _ := json.MarshalIndent(response, "", "  ")
		fmt.Println(string(jsonData))
	}
}