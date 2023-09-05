package manager

import "clean-code/repository"

type RepoManager interface {
	UomRepo() repository.UomRepository
	ProductRepo() repository.ProductRepository
	EmployeeRepo() repository.EmployeeRepository
	CustomerRepo() repository.CustomerRepository
	BillRepo() repository.BillRepository
}

type repoManager struct {
	infraManager InfraManager
}

// BillRepo implements RepoManager.
func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infraManager.Conn())
}

// CustomerRepo implements RepoManager.
func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infraManager.Conn())
}

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infraManager.Conn())
}

// ProductRepo implements RepoManager.
func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infraManager.Conn())
}

// UomRepo implements RepoManager.
func (r *repoManager) UomRepo() repository.UomRepository {
	return repository.NewUomRepository(r.infraManager.Conn())
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
