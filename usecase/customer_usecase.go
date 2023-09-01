package usecase

import (
	"clean-code/model"
	"clean-code/repository"
	"fmt"
)

type CustomerUseCase interface {
	CreateNew(payload model.Customer) error
	GetById(id string) (model.Customer, error)
	GetByPhone(phone string) (model.Customer, error)
	GetAll() ([]model.Customer, error)
	Update(payload model.Customer) error
	Delete(id string) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

// GetByPhone implements CustomerUseCase.
func (c *customerUseCase) GetByPhone(phone string) (model.Customer, error) {
	c2, err := c.repo.FindByPhone(phone)
	if err != nil {
		return c2, fmt.Errorf("failed to find phone number: %v", err)
	}

	return c2, nil
}

// CreateNew implements CustomerUseCase.
func (c *customerUseCase) CreateNew(payload model.Customer) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}

	c2, _ := c.repo.FindByPhone(payload.PhoneNumber)
	if payload.PhoneNumber == c2.PhoneNumber {
		return fmt.Errorf("phone number has registered")
	}

	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}

	err := c.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new customer: %v", err)
	}

	return nil
}

// Delete implements CustomerUseCase.
func (c *customerUseCase) Delete(id string) error {
	cust, err := c.GetById(id)
	if err != nil {
		return err
	}

	err = c.repo.DeleteById(cust.ID)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %v", err)
	}

	return nil
}

// GetAll implements CustomerUseCase.
func (c *customerUseCase) GetAll() ([]model.Customer, error) {
	c2, err := c.repo.FindByAll()
	if err != nil {
		return c2, fmt.Errorf("data is empty")
	}

	return c2, nil
}

// GetById implements CustomerUseCase.
func (c *customerUseCase) GetById(id string) (model.Customer, error) {
	c2, err := c.repo.FindById(id)
	if err != nil {
		return c2, fmt.Errorf("failed to find customer: %v", err)
	}

	return c2, nil
}

// Update implements CustomerUseCase.
func (c *customerUseCase) Update(payload model.Customer) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}

	c2, _ := c.repo.FindByPhone(payload.PhoneNumber)
	if payload.PhoneNumber == c2.PhoneNumber {
		return fmt.Errorf("phone number has registered")
	}

	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}

	_, err := c.repo.FindById(payload.ID)
	if err != nil {
		return fmt.Errorf("id is not found")
	}

	err = c.repo.UpdateById(payload)
	if err != nil {
		return fmt.Errorf("failed to update customer: %v", err)
	}

	return nil
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
