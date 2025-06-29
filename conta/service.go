package conta

import "fmt"

type ContaService struct {
	repo *Repositorio
}

func NovoContaService(repo *Repositorio) *ContaService {
	return &ContaService{
		repo: repo,
	}
}

func (s *ContaService) CriarConta(cliente, numero, agencia, chavePix string) (*Conta, error) {
	_, err := s.repo.BuscarConta(numero)
	if err == nil {
		return nil, fmt.Errorf("conta %s já existe", numero)
	}
	novaconta := NovaConta(cliente, numero, agencia, chavePix)
	s.repo.Adicionar(novaconta)
	return novaconta,nil
}

func (s *ContaService) Depositar(numeroConta string, valor float64) error {
	c, err := s.repo.BuscarConta(numeroConta)
	if err != nil {
		return err
	}
	c.Depositar(valor)
	return nil
}

func (s *ContaService) FazerPix(chavePix string, valor float64) error {
	c, err := s.repo.BuscarContaPix(chavePix)
	if err != nil {
		return err
	}
	c.FazerPix(valor)
	return nil
}

func (s *ContaService) Sacar(numeroConta string, valor float64) error {
	c, err := s.repo.BuscarConta(numeroConta)
	if err != nil {
		return err
	}
	return c.Sacar(valor)
}

func (s *ContaService) ConsultarConta(numeroConta string) (*Conta, error) {
	return s.repo.BuscarConta(numeroConta)
}

func (s *ContaService) ConsultarContaPorPix(chavePix string) (*Conta, error) {
	return s.repo.BuscarContaPix(chavePix)
}
