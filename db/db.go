package db

import "errors"

type Repository struct {
	products []*Product
}

func New() *Repository {
	return &Repository{
		products: []*Product{},
	}
}

func (r *Repository) AddProduct(p *Product) error {
	if p == nil {
		return errors.New("product is nil")
	}
	if p.Title == "" {
		return errors.New("title is empty")
	}
	id := int64(1)
	if len(r.products) > 0 {
		lastProduct := r.products[len(r.products)]
		id = lastProduct.ID + 1
	}

	p.ID = id
	r.products = append(r.products, p)
	return nil
}

func (r *Repository) GetProducts() []*Product {
	return r.products
}