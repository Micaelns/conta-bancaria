package mocks

import (
	"conta-bancaria/models"
	"errors"
)

type ContaRepoMock struct {
	ContaFake *models.Conta
}

func (m *ContaRepoMock) BuscarConta(numero string) (*models.Conta, error) {
	if m.ContaFake == nil {
		return nil, errors.New("conta não existe")
	}
	return m.ContaFake, nil
}

func (m *ContaRepoMock) BuscarContaPix(chavePix string) (*models.Conta, error) {
	if m.ContaFake == nil {
		return nil, errors.New("conta não existe")
	}
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