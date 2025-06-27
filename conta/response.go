package conta

type OperacaoResponse struct {
	Tipo     string  `json:"tipo"`
	Valor    float64 `json:"valor"`
	DataHora string  `json:"data_hora"`
	Status   string  `json:"status"`
}

type ContaResponse struct {
	Cliente   string             `json:"cliente"`
	Numero    string             `json:"numero"`
	Agencia   string             `json:"agencia"`
	ChavePix  string             `json:"chave_pix"`
	Saldo     float64            `json:"saldo"`
	Operacoes []OperacaoResponse `json:"operacoes"`
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
			Valor:    op.Valor,
			Status:   status,
		})
	}

	return ContaResponse{
		Cliente:   c.GetCliente(),
		Numero:    c.GetNumero(),
		Agencia:   c.GetAgencia(),
		ChavePix:  c.GetChavePix(),
		Saldo:     c.GetSaldo(),
		Operacoes: operacoes,
	}
}
