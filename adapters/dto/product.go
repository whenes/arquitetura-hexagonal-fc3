package dto

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	p.Name = product.Name
	p.Price = product.Price
	p.Status = product.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
}