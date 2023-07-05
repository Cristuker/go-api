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

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}
