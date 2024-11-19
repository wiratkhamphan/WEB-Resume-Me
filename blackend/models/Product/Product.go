package models

import "github.com/gofiber/fiber/v2"

// Product represents the product entity
type Product struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Qty    int    `json:"qty"`
	Detail string `json:"detail"`
}

// ProductService defines the interface for product operations
type ProductService interface {
	Create(product Product) error
	GetByID(id int) (*Product, error)
	Update(id int, product Product) error
	Delete(id int) error
	List() ([]Product, error)
}

// ProductServiceImpl is a concrete implementation of ProductService
type ProductServiceImpl struct {
	products []Product
}

// NewProductService creates a new instance of ProductServiceImpl
func NewProductService() *ProductServiceImpl {
	return &ProductServiceImpl{
		products: []Product{},
	}
}

// Create adds a new product
func (p *ProductServiceImpl) Create(product Product) error {
	product.Id = len(p.products) + 1 // Simulating auto-increment
	p.products = append(p.products, product)
	return nil
}

// GetByID retrieves a product by its ID
func (p *ProductServiceImpl) GetByID(id int) (*Product, error) {
	for _, prod := range p.products {
		if prod.Id == id {
			return &prod, nil
		}
	}
	return nil, fiber.NewError(fiber.StatusNotFound, "Product not found")
}

// Update modifies an existing product
func (p *ProductServiceImpl) Update(id int, updatedProduct Product) error {
	for i, prod := range p.products {
		if prod.Id == id {
			p.products[i] = updatedProduct
			p.products[i].Id = id // Retain original ID
			return nil
		}
	}
	return fiber.NewError(fiber.StatusNotFound, "Product not found")
}

// Delete removes a product by its ID
func (p *ProductServiceImpl) Delete(id int) error {
	for i, prod := range p.products {
		if prod.Id == id {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return nil
		}
	}
	return fiber.NewError(fiber.StatusNotFound, "Product not found")
}

// List retrieves all products
func (p *ProductServiceImpl) List() ([]Product, error) {
	return p.products, nil
}
