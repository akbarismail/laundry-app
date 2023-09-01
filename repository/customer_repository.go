package repository

import (
	"clean-code/model"
	"database/sql"
)

type CustomerRepository interface {
	Save(cust model.Customer) error
	FindById(id string) (model.Customer, error)
	FindByPhone(phone string) (model.Customer, error)
	FindByAll() ([]model.Customer, error)
	UpdateById(cust model.Customer) error
	DeleteById(id string) error
}

type customerRepository struct {
	db *sql.DB
}

// FindByPhone implements CustomerRepository.
func (c *customerRepository) FindByPhone(phone string) (model.Customer, error) {
	r := c.db.QueryRow("SELECT id, name, phone_number, address FROM customer WHERE phone_number=$1;", phone)

	cust := model.Customer{}
	err := r.Scan(&cust.ID, &cust.Name, &cust.PhoneNumber, &cust.Address)
	if err != nil {
		return cust, err
	}

	return cust, nil
}

// DeleteById implements CustomerRepository.
func (c *customerRepository) DeleteById(id string) error {
	_, err := c.db.Exec("DELETE FROM customer WHERE id=$1;", id)
	if err != nil {
		return err
	}

	return nil
}

// FindByAll implements CustomerRepository.
func (c *customerRepository) FindByAll() ([]model.Customer, error) {
	r, err := c.db.Query("SELECT id, name, phone_number, address FROM customer;")
	if err != nil {
		return nil, err
	}

	defer r.Close()

	customers := []model.Customer{}
	for r.Next() {
		customer := model.Customer{}
		err = r.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

// FindById implements CustomerRepository.
func (c *customerRepository) FindById(id string) (model.Customer, error) {
	r := c.db.QueryRow("SELECT id, name, phone_number, address FROM customer WHERE id=$1;", id)

	cust := model.Customer{}
	err := r.Scan(&cust.ID, &cust.Name, &cust.PhoneNumber, &cust.Address)
	if err != nil {
		return cust, err
	}

	return cust, nil
}

// Save implements CustomerRepository.
func (c *customerRepository) Save(cust model.Customer) error {
	_, err := c.db.Exec("INSERT INTO customer VALUES ($1, $2, $3, $4);", cust.ID, cust.Name, cust.PhoneNumber, cust.Address)
	if err != nil {
		return err
	}

	return nil
}

// UpdateById implements CustomerRepository.
func (c *customerRepository) UpdateById(cust model.Customer) error {
	_, err := c.db.Exec("UPDATE customer SET name=$1, phone_number=$2, address=$3  WHERE id=$4;", cust.Name, cust.PhoneNumber, cust.Address, cust.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
