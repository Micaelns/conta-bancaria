package conta

import "time"

type TipoOperacao string

const (
	Deposito TipoOperacao = "Depósito"
	Pix    TipoOperacao = "Pix"
	Saque    TipoOperacao = "Saque"
)

type Operacao struct {
	Tipo      TipoOperacao
	Valor     float64
	Erro      string
	DataHora time.Time
}
