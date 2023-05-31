package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

/*
Aqui eu criei um serviço de get sem sequer conhecer o banco, dessa forma eu isolo a minha aplicação eu trabalho com adapters para conectar ao banco.
Como consequência disso a minha regra de negócio não sofrera interferência dos serviços que eu for usar
*/
func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}
	return product, nil
}
