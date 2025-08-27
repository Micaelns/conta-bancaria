package models

import (
	"testing"
)

func TestContaDepositar(t *testing.T) {
	t.Run("Sucesso", func(t *testing.T) {
		conta := &Conta{Saldo: 0} 

		op := conta.Depositar(100)
		if op.Erro != "" {
			t.Fatalf("esperava sucesso, mas ocorreu erro: %v", op.Erro)
		}

		if conta.Saldo != 100 {
			t.Errorf("esperava saldo 100, mas obteve %.2f", conta.Saldo)
		}

		if op.Tipo != string(Deposito) {
			t.Errorf("esperava 1 operação Deposito, mas obteve %v", op.Tipo)
		}
	})

	t.Run("Valor inválido", func(t *testing.T) {
		conta := &Conta{Saldo: 0}
		 
		op := conta.Depositar(-10)
		if op.Erro == "" {
			t.Fatal("esperava erro ao depositar valor inválido, mas não ocorreu")
		}

		if conta.Saldo != 0 {
			t.Errorf("saldo não deveria ter mudado, mas obteve %.2f", conta.Saldo)
		}
	})
}

func TestContaFazerPix(t *testing.T) {
	t.Run("Sucesso", func(t *testing.T) {
		conta := &Conta{Saldo: 0}
		
		op := conta.FazerPix(100)
		if op.Erro != "" {
			t.Fatalf("esperava sucesso, mas ocorreu erro: %v", op.Erro)
		}
		
		if conta.Saldo != 100 {
			t.Errorf("esperava saldo 100, mas obteve %.2f", conta.Saldo)
		}

		if op.Tipo != string(Pix) {
			t.Errorf("esperava 1 operação Pix, mas obteve %v", op.Tipo)
		}
	})
 
	t.Run("Valor inválido", func(t *testing.T) {
		conta := &Conta{Saldo: 0}
		
		op := conta.FazerPix(-10)
		if op.Erro == "" {
			t.Fatal("esperava erro ao realizar pix com valor inválido, mas não ocorreu")
		}

		if conta.Saldo != 0 {
			t.Errorf("saldo não deveria ter mudado, mas obteve %.2f", conta.Saldo)
		}
	})
}