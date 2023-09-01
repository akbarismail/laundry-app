package usecase

import (
	"clean-code/model"
	"clean-code/repository"
	"fmt"
)

type BillUseCase interface {
	CreateNew(payload model.Bill) error
	GetAll() ([]model.Bill, error)
	GetById(id string) (model.Bill, error)
}

type billUseCase struct {
	repo       repository.BillRepository
	employeeUC EmployeeUseCase
	customerUC CustomerUseCase
	productUC  ProductUseCase
}

// GetById implements BillUseCase.
func (b *billUseCase) GetById(id string) (model.Bill, error) {
	panic("unimplemented")
}

// GetAll implements BillUseCase.
func (b *billUseCase) GetAll() ([]model.Bill, error) {
	panic("unimplemented")
}

// CreateNew implements BillUseCase.
func (b *billUseCase) CreateNew(payload model.Bill) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}

	_, err := b.employeeUC.GetById(payload.Employee.ID)
	if err != nil {
		return fmt.Errorf("employee is not found")
	}

	_, err = b.customerUC.GetById(payload.Customer.ID)
	if err != nil {
		return fmt.Errorf("customer is not found")
	}

	for _, billDetail := range payload.BillDetail {
		_, err = b.productUC.GetById(billDetail.Product.ID)
		if err != nil {
			return fmt.Errorf("product is not found")
		}
	}

	err = b.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create bill: %v", err)
	}

	return nil
}

func NewBillUseCase(repo repository.BillRepository, employeeUC EmployeeUseCase, customerUC CustomerUseCase, productUC ProductUseCase) BillUseCase {
	return &billUseCase{
		repo:       repo,
		employeeUC: employeeUC,
		customerUC: customerUC,
		productUC:  productUC,
	}
}
