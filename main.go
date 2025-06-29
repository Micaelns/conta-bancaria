package main

import (
	"conta-bancaria/api/controllers"
	"conta-bancaria/api/routes"
	"conta-bancaria/conta"
)

func main() {
	repo := conta.NovoRepositorio()
	service := conta.NovoContaService(repo)
	controller := controllers.NovoContaController(service)

	// Criar contas iniciais para teste
	conta1, _ := service.CriarConta("João Silva", "123", "0001", "joao@email.com")
	conta1.Depositar(400)
	conta1.Depositar(400)
	conta1.Sacar(300)
	conta1.Sacar(900)

	r := routes.SetupRouter(controller)
	r.Run(":8080")
}
