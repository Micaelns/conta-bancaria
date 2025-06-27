package conta

type ContaService struct {
	repo *Repositorio
}

func NovoContaService(repo *Repositorio) *ContaService {
	return &ContaService{
		repo: repo,
	}
}

func (s *ContaService) CriarConta(cliente, numero, agencia, chavePix string) *Conta {
	c := NovaConta(cliente, numero, agencia, chavePix)
	s.repo.Adicionar(c)
	return c
}

func (s *ContaService) Depositar(numeroConta string, valor float64) error {
	c, err := s.repo.BuscarConta(numeroConta)
	if err != nil {
		return err
	}
	c.Depositar(valor)
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
