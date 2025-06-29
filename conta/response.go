package conta

import "fmt"

type OperacaoResponse struct {
	Tipo     string `json:"tipo"`
	Valor    string `json:"valor"`
	DataHora string `json:"data_hora"`
	Status   string `json:"status"`
}

type ContaResponse struct {
	Cliente   string             `json:"cliente"`
	Numero    string             `json:"numero"`
	Agencia   string             `json:"agencia"`
	ChavePix  string             `json:"chave_pix"`
	Saldo     string             `json:"saldo"`
	Operacoes []OperacaoResponse `json:"operacoes"`
}
type ContaSaldoResponse struct {
	Cliente   string             `json:"cliente"`
	Numero    string             `json:"numero"`
	Agencia   string             `json:"agencia"`
	ChavePix  string             `json:"chave_pix"`
	Saldo     string             `json:"saldo"`
}

func (c *Conta) ToSaldoResponse() ContaSaldoResponse {
	return ContaSaldoResponse{
		Cliente:   c.GetCliente(),
		Numero:    c.GetNumero(),
		Agencia:   c.GetAgencia(),
		ChavePix:  c.GetChavePix(),
		Saldo:     fmt.Sprintf("R$ %.2f", c.GetSaldo()),
	}
}

func (c *Conta) ToResponse() ContaResponse {
	var operacoes []OperacaoResponse

	for _, op := range c.operacoes {
		status := "OK"
		if op.Erro != "" {
			status = "FALHA: " + op.Erro
		}

		operacoes = append(operacoes, OperacaoResponse{
			DataHora: op.DataHora.Format("02/01/2006 15:04:05"),
			Tipo:     string(op.Tipo),
			Valor:    fmt.Sprintf("R$ %.2f", op.Valor),
			Status:   status,
		})
	}

	return ContaResponse{
		Cliente:   c.GetCliente(),
		Numero:    c.GetNumero(),
		Agencia:   c.GetAgencia(),
		ChavePix:  c.GetChavePix(),
		Saldo:     fmt.Sprintf("R$ %.2f", c.GetSaldo()),
		Operacoes: operacoes,
	}
}
