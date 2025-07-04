package main

import (
	"conta-bancaria/api/controllers"
	"conta-bancaria/api/routes"
	"conta-bancaria/infra"
	"conta-bancaria/repositories"
	"conta-bancaria/services"
	"log"
)

func main() {
	
	if err := infra.ConectarMongo(); err != nil {
		log.Fatal("Erro ao conectar ao MongoDB:", err)
	} 

	contaRepo := repositories.NovaContaRepository()
	operacaoRepo := repositories.NovaOperacaoRepository()
	service := services.NovoContaService(contaRepo, operacaoRepo)
	 
	controller := controllers.NovoContaController(service)
	r := routes.SetupRouter(controller)
	r.Run(":8080") 
}