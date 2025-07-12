package services_test

import (
	"conta-bancaria/models"
	"conta-bancaria/services"
	"conta-bancaria/services/mocks"
	"testing"
)

func TestDepositar(t *testing.T) {
	conta := &models.Conta{
		Numero: "123",
		Saldo:  0,
	}

	mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
	mockOperacaoRepo := &mocks.OperacaoRepoMock{}

	service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

	err := service.Depositar("123", 100)
	if err != nil {
		t.Fatalf("esperava sucesso, mas ocorreu erro: %v", err)
	}

	if conta.Saldo != 100 {
		t.Errorf("esperava saldo 100, mas obteve %.2f", conta.Saldo)
	}

	if len(mockOperacaoRepo.Registradas) != 1 {
		t.Errorf("esperava 1 operação registrada, mas obteve %d", len(mockOperacaoRepo.Registradas))
	}
}

func TestFazerPix(t *testing.T) {
	conta := &models.Conta{
		ChavePix: "pix@gmail.com",
		Saldo:  0,
	}

	mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
	mockOperacaoRepo := &mocks.OperacaoRepoMock{}

	service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

	err := service.FazerPix("pix@gmail.com", 100)
	if err != nil {
		t.Fatalf("esperava sucesso, mas ocorreu erro: %v", err)
	}

	if conta.Saldo != 100 {
		t.Errorf("esperava saldo 100, mas obteve %.2f", conta.Saldo)
	}

	if len(mockOperacaoRepo.Registradas) != 1 {
		t.Errorf("esperava 1 operação registrada, mas obteve %d", len(mockOperacaoRepo.Registradas))
	}
}

func TestSacar(t *testing.T) {
	conta := &models.Conta{
		Numero: "123",
		Saldo:  0,
	}

	mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
	mockOperacaoRepo := &mocks.OperacaoRepoMock{}

	service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

	service.Depositar("123", 300)
	err := service.Sacar("123", 100)
	if err != nil {
		t.Fatalf("esperava sucesso, mas ocorreu erro: %v", err)
	}

	if conta.Saldo != 200 {
		t.Errorf("esperava saldo 200, mas obteve %.2f", conta.Saldo)
	}

	if len(mockOperacaoRepo.Registradas) != 2 {
		t.Errorf("esperava 2 operação registrada, mas obteve %d", len(mockOperacaoRepo.Registradas))
	}
}
 