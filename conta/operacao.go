package conta

import "time"

type TipoOperacao string

const (
	Deposito TipoOperacao = "Depósito"
	Saque    TipoOperacao = "Saque"
)

type Operacao struct {
	Tipo      TipoOperacao
	Valor     float64
	Erro      string
	DataHora time.Time
}
