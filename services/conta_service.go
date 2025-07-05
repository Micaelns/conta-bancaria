package services

import (
	"conta-bancaria/models"
	"conta-bancaria/repositories"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type ContaService struct {
	contaRepo *repositories.ContaRepository
	operacaoRepo *repositories.OperacaoRepository
}

func NovoContaService(contaRepo *repositories.ContaRepository, operacaoRepo *repositories.OperacaoRepository) *ContaService {
	return &ContaService{
		contaRepo: contaRepo,
		operacaoRepo: operacaoRepo,
	}
}

func (s *ContaService) CriarConta(cliente, numero, agencia, chavePix string) (*models.Conta, error) {
	_, err := s.contaRepo.BuscarConta(numero)
	if err == nil {
		return nil, errors.New("conta já existe")
	}
	novaconta := gerarConta(cliente, numero, agencia, chavePix)
	
	err2 := s.contaRepo.Adicionar(novaconta)
	if err2 != nil {
		return nil, err2
	}
	return novaconta,nil
}

func gerarConta(cliente string, numero string, agencia string, chavePix string) *models.Conta{
	return &models.Conta{
		Cliente: cliente,
		Numero: numero,
		Agencia: agencia,
		ChavePix: chavePix,
		Saldo: 0,
	}
}

func (s *ContaService) Depositar(numeroConta string, valor float64) error {
	c, err := s.contaRepo.BuscarConta(numeroConta)
	if err != nil {
		return err
	}
	operacao := c.Depositar(valor)
	s.operacaoRepo.Registrar(operacao)
	if operacao.Erro != "" {
		return fmt.Errorf(operacao.Erro)
	}
	s.contaRepo.AtualizarSaldo(numeroConta, c.Saldo)
	return nil
}

func (s *ContaService) FazerPix(chavePix string, valor float64) error {
	c, err := s.contaRepo.BuscarContaPix(chavePix)
	if err != nil {
		return err
	}
	operacao := c.FazerPix(valor)
	s.operacaoRepo.Registrar(operacao)
	if operacao.Erro != "" {
		return fmt.Errorf(operacao.Erro)
	}
	s.contaRepo.AtualizarSaldo(c.Numero, c.Saldo)
	return nil
}

func (s *ContaService) Sacar(numeroConta string, valor float64) error {
	c, err := s.contaRepo.BuscarConta(numeroConta)
	if err != nil {
		return err
	}
	operacao := c.Sacar(valor)
	s.operacaoRepo.Registrar(operacao)
	if operacao.Erro != "" {
		return fmt.Errorf(operacao.Erro)
	}
	s.contaRepo.AtualizarSaldo(c.Numero, c.Saldo)
	return nil
}

func (s *ContaService) ConsultarConta(numeroConta string) (*models.Conta, error) {
	return s.contaRepo.BuscarConta(numeroConta)
}

func (s *ContaService) ConsultarContaPorPix(chavePix string) (*models.Conta, error) {
	return s.contaRepo.BuscarContaPix(chavePix)
}

func (s *ContaService) ExtratoSimples(numeroConta string) ([]models.Operacao, error) {
	
	filter := bson.M{
		"numeroConta": numeroConta,
		"$or": []bson.M{
			{"erro": bson.M{"$eq": ""}},
			{"erro": bson.M{"$exists": false}},
		},
	}

	return s.operacaoRepo.ListaFiltrada(filter)
}

func (s *ContaService) ExtratoCompleto(numeroConta string) ([]models.Operacao, error) {
	filter := bson.M{ "numeroConta": numeroConta }
	return s.operacaoRepo.ListaFiltrada(filter)
}