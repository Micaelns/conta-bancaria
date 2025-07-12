package mocks

import (
	"conta-bancaria/models"
)

type ContaRepoMock struct {
	ContaFake *models.Conta
}

func (m *ContaRepoMock) BuscarConta(numero string) (*models.Conta, error) {
	return m.ContaFake, nil
}

func (m *ContaRepoMock) BuscarContaPix(chavePix string) (*models.Conta, error) {
	return m.ContaFake, nil
}

func (m *ContaRepoMock) Adicionar(conta *models.Conta) error {
	m.ContaFake = conta
	return nil
}

func (m *ContaRepoMock) AtualizarSaldo(numero string, novoSaldo float64) error {
	m.ContaFake.Saldo = novoSaldo
	return nil
}