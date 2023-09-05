package manager

import "clean-code/usecase"

type UseCaseManager interface {
	BillUC() usecase.BillUseCase
	CustomerUC() usecase.CustomerUseCase
	EmployeeUC() usecase.EmployeeUseCase
	ProductUC() usecase.ProductUseCase
	UomUC() usecase.UomUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// BillUC implements UseCaseManager.
func (u *useCaseManager) BillUC() usecase.BillUseCase {
	return usecase.NewBillUseCase(u.repoManager.BillRepo(), u.EmployeeUC(), u.CustomerUC(), u.ProductUC())
}

// CustomerUC implements UseCaseManager.
func (u *useCaseManager) CustomerUC() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

// EmployeeUC implements UseCaseManager.
func (u *useCaseManager) EmployeeUC() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

// ProductUC implements UseCaseManager.
func (u *useCaseManager) ProductUC() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManager.ProductRepo(), u.UomUC())
}

// UomUC implements UseCaseManager.
func (u *useCaseManager) UomUC() usecase.UomUseCase {
	return usecase.NewUomUseCase(u.repoManager.UomRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
