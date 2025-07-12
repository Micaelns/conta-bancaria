package routes

import (
	"conta-bancaria/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(contaController *controllers.ContaController) *gin.Engine {
	r := gin.Default()

	r.POST("/contas", contaController.CriarConta)
	r.POST("/contas/:numero/depositar", contaController.Depositar)
	r.POST("/contas/:numero/sacar", contaController.Sacar)
	r.GET("/contas/:numero/saldo", contaController.Saldo)
	r.GET("/contas/:numero/extrato", contaController.Extrato)

	r.POST("/pix/:chavePix/fazerPix", contaController.FazerPix)

	return r
}
