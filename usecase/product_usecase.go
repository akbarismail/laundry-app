package usecase

import (
	"clean-code/model"
	"clean-code/repository"
	"fmt"
)

type ProductUseCase interface {
	CreateNew(payload model.Product) error
	GetById(id string) (model.Product, error)
	GetByName(name string) ([]model.Product, error)
	GetAll() ([]model.Product, error)
	Update(payload model.Product) error
	Delete(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
	uom  UomUseCase
}

// GetByName implements ProductUseCase.
func (p *productUseCase) GetByName(name string) ([]model.Product, error) {
	p2, err := p.repo.FindByName(name)
	if err != nil {
		return p2, err
	}

	return p2, nil
}

// CreateNew implements ProductUseCase.
func (p *productUseCase) CreateNew(payload model.Product) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price <= 0 {
		return fmt.Errorf("price is must be greater than zero")
	}

	_, err := p.uom.GetById(payload.Uom.ID)
	if err != nil {
		return fmt.Errorf("uom is not found")
	}

	err = p.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new product: %v", err)
	}

	return nil
}

// Delete implements ProductUseCase.
func (p *productUseCase) Delete(id string) error {
	product, err := p.GetById(id)
	if err != nil {
		return err
	}

	err = p.repo.DeleteById(product.ID)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	return nil
}

// GetAll implements ProductUseCase.
func (p *productUseCase) GetAll() ([]model.Product, error) {
	p2, err := p.repo.FindByAll()
	if err != nil {
		return p2, fmt.Errorf("data product is empty")
	}
	return p2, nil
}

// GetById implements ProductUseCase.
func (p *productUseCase) GetById(id string) (model.Product, error) {
	p2, err := p.repo.FindById(id)
	if err != nil {
		return model.Product{}, fmt.Errorf("failed to find product: %v", err)
	}

	return p2, nil
}

// Update implements ProductUseCase.
func (p *productUseCase) Update(payload model.Product) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price <= 0 {
		return fmt.Errorf("price is must be greater than zero")
	}

	_, err := p.uom.GetById(payload.Uom.ID)
	if err != nil {
		return fmt.Errorf("id uom is not found")
	}

	err = p.repo.UpdateById(payload)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}

	return nil
}

func NewProductUseCase(repo repository.ProductRepository, uom UomUseCase) ProductUseCase {
	return &productUseCase{
		repo: repo,
		uom:  uom,
	}
}
