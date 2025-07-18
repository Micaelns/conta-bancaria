package controllers

import (
	"conta-bancaria/models"
	"conta-bancaria/services"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NovaContaRequest struct {
	Cliente  string `json:"cliente" binding:"required"`
	Numero   string `json:"numero" binding:"required"`
	Agencia  string `json:"agencia" binding:"required"`
	ChavePix string `json:"chavePix" binding:"required"`
}

type OperacaoRequest struct {
	Valor float64 `json:"valor" binding:"required"`
}

type ContaController struct {
	Service *services.ContaService
}

func NovoContaController(service *services.ContaService) *ContaController {
	return &ContaController{Service: service}
}

func (cc *ContaController) CriarConta(c *gin.Context) {
	var req NovaContaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	conta, err := cc.Service.CriarConta(req.Cliente, req.Numero, req.Agencia, req.ChavePix)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, conta)
}

func (cc *ContaController) Depositar(c *gin.Context) {
	numero := c.Param("numero")
	var req OperacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := cc.Service.Depositar(numero, req.Valor)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	conta, _ := cc.Service.ConsultarConta(numero)
	c.JSON(http.StatusOK, conta)
}

func (cc *ContaController) FazerPix(c *gin.Context) {
	chavePix := c.Param("chavePix")
	var req OperacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := cc.Service.FazerPix(chavePix, req.Valor)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	conta, _ := cc.Service.ConsultarContaPorPix(chavePix)
	c.JSON(http.StatusOK, conta)
}

func (cc *ContaController) Sacar(c *gin.Context) {
	numero := c.Param("numero")
	var req OperacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := cc.Service.Sacar(numero, req.Valor)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": err.Error()})
		return
	}

	conta, _ := cc.Service.ConsultarConta(numero)
	c.JSON(http.StatusOK, conta)
}

func (cc *ContaController) Saldo(c *gin.Context) {
	numero := c.Param("numero")

	conta, err := cc.Service.ConsultarConta(numero)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	conta.Saldo = math.Round(conta.Saldo*100) / 100
	c.JSON(http.StatusOK, conta)
}

func (cc *ContaController) Extrato(c *gin.Context) {
	numero := c.Param("numero")
	completo := c.DefaultQuery("completo", "false")

	var (
		operacoes []models.Operacao
		err       error
	)
	
	if completo == "true" {
		operacoes, err = cc.Service.ExtratoCompleto(numero)
	} else {
		operacoes, err = cc.Service.ExtratoSimples(numero)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, operacoes)
}
