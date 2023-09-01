package usecase

import (
	"clean-code/model"
	"clean-code/repository"
	"fmt"
)

type EmployeeUseCase interface {
	CreateNew(payload model.Employee) error
	GetById(id string) (model.Employee, error)
	GetByPhone(phone string) (model.Employee, error)
	GetAll() ([]model.Employee, error)
	Update(payload model.Employee) error
	Delete(id string) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

// GetByPhone implements EmployeeUseCase.
func (e *employeeUseCase) GetByPhone(phone string) (model.Employee, error) {
	e2, err := e.repo.FindByPhone(phone)
	if err != nil {
		return e2, fmt.Errorf("failed to find phone number: %v", err)
	}

	return e2, nil
}

// CreateNew implements EmployeeUseCase.
func (e *employeeUseCase) CreateNew(payload model.Employee) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}

	e2, _ := e.repo.FindByPhone(payload.PhoneNumber)
	if payload.PhoneNumber == e2.PhoneNumber {
		return fmt.Errorf("phone number has registered")
	}

	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}

	err := e.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new employee: %v", err)
	}

	return nil
}

// Delete implements EmployeeUseCase.
func (e *employeeUseCase) Delete(id string) error {
	employee, err := e.GetById(id)
	if err != nil {
		return err
	}

	err = e.repo.DeleteById(employee.ID)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %v", err)
	}

	return nil
}

// GetAll implements EmployeeUseCase.
func (e *employeeUseCase) GetAll() ([]model.Employee, error) {
	e2, err := e.repo.FindByAll()
	if err != nil {
		return e2, fmt.Errorf("data is empty")
	}

	return e2, nil
}

// GetById implements EmployeeUseCase.
func (e *employeeUseCase) GetById(id string) (model.Employee, error) {
	e2, err := e.repo.FindById(id)
	if err != nil {
		return e2, fmt.Errorf("failed to find customer: %v", err)
	}

	return e2, nil
}

// Update implements EmployeeUseCase.
func (e *employeeUseCase) Update(payload model.Employee) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}

	e2, _ := e.repo.FindByPhone(payload.PhoneNumber)
	if payload.PhoneNumber == e2.PhoneNumber {
		return fmt.Errorf("phone number has registered")
	}

	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}

	_, err := e.repo.FindById(payload.ID)
	if err != nil {
		return fmt.Errorf("id is not found")
	}

	err = e.repo.UpdateById(payload)
	if err != nil {
		return fmt.Errorf("failed to update employee: %v", err)
	}

	return nil
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
