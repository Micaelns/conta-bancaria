package conta

import "fmt"

type Repositorio struct {
	contas map[string]*Conta
}

func NovoRepositorio() *Repositorio {
	return &Repositorio{
		contas: make(map[string]*Conta),
	}
}

func (r *Repositorio) Adicionar(c *Conta) {
	r.contas[c.numero] = c
}

func (r *Repositorio) BuscarConta(numero string) (*Conta, error) {
	conta, existe := r.contas[numero]
	if !existe {
		return nil, fmt.Errorf("conta %s não encontrada", numero)
	}
	return conta, nil
}
