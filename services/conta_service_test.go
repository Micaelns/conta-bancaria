package services_test

import (
	"conta-bancaria/models"
	"conta-bancaria/services"
	"conta-bancaria/services/mocks"
	"testing"
)

func TestContaServiceDepositar(t *testing.T) {
	t.Run("Sucesso", func(t *testing.T) {
		conta := &models.Conta{Numero: "123", Saldo: 0}
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
	})

	t.Run("Conta inexistente", func(t *testing.T) {
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: nil} // Simula conta não encontrada
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.Depositar("999", 50)
		if err == nil {
			t.Fatal("esperava erro ao depositar em conta inexistente, mas não ocorreu")
		}
	})

	t.Run("Valor inválido", func(t *testing.T) {
		conta := &models.Conta{Numero: "123", Saldo: 0}
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.Depositar("123", -10)
		if err == nil {
			t.Fatal("esperava erro ao realizar pix com valor inválido, mas não ocorreu")
		}

		if conta.Saldo != 0 {
			t.Errorf("saldo não deveria ter mudado, mas obteve %.2f", conta.Saldo)
		}
	})
}

func TestContaServiceFazerPix(t *testing.T) {
	t.Run("Sucesso", func(t *testing.T) {
		conta := &models.Conta{ChavePix: "pix@gmail.com", Saldo: 0}
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
	})

	t.Run("Conta inexistente", func(t *testing.T) {
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: nil} // Simula conta não encontrada
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.FazerPix("pix@999", 50)
		if err == nil {
			t.Fatal("esperava erro ao depositar em conta inexistente, mas não ocorreu")
		}
	})

	t.Run("Valor inválido", func(t *testing.T) {
		conta := &models.Conta{ChavePix: "pix@gmail.com", Saldo: 0}
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.FazerPix("pix@gmail.com", -10)
		if err == nil {
			t.Fatal("esperava erro ao depositar valor inválido, mas não ocorreu")
		}

		if conta.Saldo != 0 {
			t.Errorf("saldo não deveria ter mudado, mas obteve %.2f", conta.Saldo)
		}
	})
}

func TestContaServiceSacar(t *testing.T) {
	t.Run("Sucesso", func(t *testing.T) {
		conta := &models.Conta{Numero: "123", Saldo: 200}
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.Sacar("123", 100)
		if err != nil {
			t.Fatalf("esperava sucesso, mas ocorreu erro: %v", err)
		}

		if conta.Saldo != 100 {
			t.Errorf("esperava saldo 100, mas obteve %.2f", conta.Saldo)
		}

		if len(mockOperacaoRepo.Registradas) != 1 {
			t.Errorf("esperava 1 operação registrada, mas obteve %d", len(mockOperacaoRepo.Registradas))
		}
	})

	t.Run("Conta inexistente", func(t *testing.T) {
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: nil} // Simula conta não encontrada
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.Sacar("999", 50)
		if err == nil {
			t.Fatal("esperava erro ao sacar em conta inexistente, mas não ocorreu")
		}
	})

	t.Run("Saldo insuficiente", func(t *testing.T) {
		conta := &models.Conta{Numero: "123", Saldo: 50}
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.Sacar("123", 100)
		if err == nil {
			t.Fatal("esperava erro ao sacar valor insuficiente do saldo, mas não ocorreu")
		}

		if conta.Saldo != 50 {
			t.Errorf("saldo não deveria ter mudado, mas obteve %.2f", conta.Saldo)
		}
	})

	t.Run("Valor inválido", func(t *testing.T) {
		conta := &models.Conta{Numero: "123", Saldo: 50}
		mockContaRepo := &mocks.ContaRepoMock{ContaFake: conta}
		mockOperacaoRepo := &mocks.OperacaoRepoMock{}
		service := services.NovoContaService(mockContaRepo, mockOperacaoRepo)

		err := service.Sacar("123", -10)
		if err == nil {
			t.Fatal("esperava erro ao sacar valor inválido, mas não ocorreu")
		}

		if conta.Saldo != 50 {
			t.Errorf("saldo não deveria ter mudado, mas obteve %.2f", conta.Saldo)
		}
	})
}
 