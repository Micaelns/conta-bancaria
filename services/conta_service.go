package services

import (
	"conta-bancaria/models"
	"conta-bancaria/repositories"
	"errors"
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
	operacao, err2 := c.Depositar(valor)
	if err2 != nil {
		return err
	}
	s.contaRepo.AtualizarSaldo(numeroConta, c.Saldo)
	s.operacaoRepo.Registrar(operacao)
	return nil
}

func (s *ContaService) FazerPix(chavePix string, valor float64) error {
	c, err := s.contaRepo.BuscarContaPix(chavePix)
	if err != nil {
		return err
	}
	operacao, err2 := c.FazerPix(valor)
	if err2 != nil {
		return err
	}
	s.contaRepo.AtualizarSaldo(c.Numero, c.Saldo)
	s.operacaoRepo.Registrar(operacao)
	return nil
}

func (s *ContaService) Sacar(numeroConta string, valor float64) error {
	c, err := s.contaRepo.BuscarConta(numeroConta)
	if err != nil {
		return err
	}
	operacao, err2 := c.Sacar(valor)
	if err2 != nil {
		return err2
	}
	s.contaRepo.AtualizarSaldo(c.Numero, c.Saldo)
	s.operacaoRepo.Registrar(operacao)
	return nil
}

func (s *ContaService) ConsultarConta(numeroConta string) (*models.Conta, error) {
	return s.contaRepo.BuscarConta(numeroConta)
}

func (s *ContaService) ConsultarContaPorPix(chavePix string) (*models.Conta, error) {
	return s.contaRepo.BuscarContaPix(chavePix)
}

func (s *ContaService) Extrato(numeroConta string) ([]models.Operacao, error) {
	return s.operacaoRepo.ListarPorConta(numeroConta)
}