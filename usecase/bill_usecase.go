package usecase

import (
	"clean-code/model"
	"clean-code/repository"
	"fmt"
)

type BillUseCase interface {
	CreateNew(payloadBill model.Bill, payloadBillDetail model.BillDetail) error
	GetAll() ([]model.Bill, []model.BillDetail, error)
	GetById(id string) (model.Bill, []model.BillDetail, error)
}

type billUseCase struct {
	repo       repository.BillRepository
	employeeUC EmployeeUseCase
	customerUC CustomerUseCase
	productUC  ProductUseCase
}

// GetById implements BillUseCase.
func (b *billUseCase) GetById(id string) (model.Bill, []model.BillDetail, error) {
	bill, billsDetail, err := b.repo.FindById(id)
	if err != nil {
		return bill, nil, fmt.Errorf("failed to find bill: %v", err)
	}

	return bill, billsDetail, nil
}

// GetAll implements BillUseCase.
func (b *billUseCase) GetAll() ([]model.Bill, []model.BillDetail, error) {
	bills, billsDetail, err := b.repo.FindByAll()
	if err != nil {
		return nil, nil, fmt.Errorf("data is empty")
	}

	return bills, billsDetail, nil
}

// CreateNew implements BillUseCase.
func (b *billUseCase) CreateNew(payloadBill model.Bill, payloadBillDetail model.BillDetail) error {
	if payloadBill.ID == "" {
		return fmt.Errorf("id is required")
	}

	if payloadBill.BillDate.String() == "" {
		return fmt.Errorf("bill date is required")
	}

	if payloadBill.EntryDate.String() == "" {
		return fmt.Errorf("entry date is required")
	}

	if payloadBill.FinishDate.String() == "" {
		return fmt.Errorf("finish date is required")
	}

	if payloadBill.Employee.ID == "" {
		return fmt.Errorf("employee id is required")
	}

	if payloadBill.Customer.ID == "" {
		return fmt.Errorf("customer id is required")
	}

	_, err := b.employeeUC.GetById(payloadBill.Employee.ID)
	if err != nil {
		return fmt.Errorf("employee is not found")
	}

	_, err = b.customerUC.GetById(payloadBill.Customer.ID)
	if err != nil {
		return fmt.Errorf("customer is not found")
	}

	if payloadBillDetail.ID == "" {
		return fmt.Errorf("bill detail id is required")
	}

	if payloadBillDetail.Bill.ID == "" {
		return fmt.Errorf("bill id is required")
	}

	// _, err = b.repo.FindById(payloadBillDetail.Bill.ID)
	// if err != nil {
	// 	return fmt.Errorf("bill is not found")
	// }

	_, err = b.productUC.GetById(payloadBillDetail.Product.ID)
	if err != nil {
		return fmt.Errorf("product is not found")
	}

	if payloadBillDetail.ProductPrice <= 0 {
		return fmt.Errorf("product price is must be greater than zero")
	}

	if payloadBillDetail.Qty <= 0 {
		return fmt.Errorf("qty is must be greater than zero")
	}

	err = b.repo.Save(payloadBill, payloadBillDetail)
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
