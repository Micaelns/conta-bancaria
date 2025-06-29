package conta

import (
	"testing"
)

func TestContaService_CriarConta(t *testing.T) {
	repo := NovoRepositorio()
	service := NovoContaService(repo)

	conta, err := service.CriarConta("João Silva", "123", "0001", "joao@email.com")
	if conta == nil {
		t.Fatal("Conta não pode ser nula")
	}

	if err != nil {
		t.Errorf("não esperava erro : %s", err)
	}

	if conta.GetCliente() != "João Silva" {
		t.Errorf("esperado: João Silva, obtido: %s", conta.GetCliente())
	}
}

func TestContaService_Depositar(t *testing.T) {
	repo := NovoRepositorio()
	service := NovoContaService(repo)

	service.CriarConta("Maria", "456", "0002", "maria@pix.com")
	err := service.Depositar("456", 1000)
	if err != nil {
		t.Errorf("não esperava erro no depósito: %v", err)
	}

	conta, _ := service.ConsultarConta("456")
	if conta.GetSaldo() != 1000 {
		t.Errorf("saldo esperado: 1000, obtido: %.2f", conta.GetSaldo())
	}
}

func TestContaService_FazerPix(t *testing.T) {
	repo := NovoRepositorio()
	service := NovoContaService(repo)

	service.CriarConta("Maria", "456", "0002", "maria@pix.com")
	err := service.FazerPix("maria@pix.com", 1000)
	if err != nil {
		t.Errorf("não esperava erro no pix: %v", err)
	}

	conta, _ := service.ConsultarConta("456")
	if conta.GetSaldo() != 1000 {
		t.Errorf("saldo esperado: 1000, obtido: %.2f", conta.GetSaldo())
	}
}

func TestContaService_Saque_Sucesso(t *testing.T) {
	repo := NovoRepositorio()
	service := NovoContaService(repo)

	service.CriarConta("João", "789", "0003", "joao@pix.com")
	service.Depositar("789", 500)
	err := service.Sacar("789", 300)
	if err != nil {
		t.Errorf("não esperava erro no saque: %v", err)
	}

	conta, _ := service.ConsultarConta("789")
	if conta.GetSaldo() != 200 {
		t.Errorf("saldo esperado: 200, obtido: %.2f", conta.GetSaldo())
	}
}

func TestContaService_Saque_ErroSaldoInsuficiente(t *testing.T) {
	repo := NovoRepositorio()
	service := NovoContaService(repo)

	service.CriarConta("Lucas", "999", "0004", "lucas@pix.com")
	err := service.Sacar("999", 100)
	if err == nil {
		t.Error("esperava erro de saldo insuficiente, mas não ocorreu")
	}
}
