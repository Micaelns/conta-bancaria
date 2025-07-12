// mocks/operacao_repo_mock.go
package mocks

import (
	"conta-bancaria/models"

	"go.mongodb.org/mongo-driver/bson"
)

type OperacaoRepoMock struct {
	Registradas []*models.Operacao 
}

func (m *OperacaoRepoMock) Registrar(op *models.Operacao) error {
	m.Registradas = append(m.Registradas, op)
	return nil
}

func (m *OperacaoRepoMock) ListaFiltrada(filter bson.M) ([]models.Operacao, error) {
	var result []models.Operacao
	for _, op := range m.Registradas {
		result = append(result, *op)
	}
	return result, nil
}
