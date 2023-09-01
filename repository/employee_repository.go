package repository

import (
	"clean-code/model"
	"database/sql"
)

type EmployeeRepository interface {
	Save(emp model.Employee) error
	FindById(id string) (model.Employee, error)
	FindByPhone(phone string) (model.Employee, error)
	FindByAll() ([]model.Employee, error)
	UpdateById(emp model.Employee) error
	DeleteById(id string) error
}

type employeeRepository struct {
	db *sql.DB
}

// FindByPhone implements EmployeeRepository.
func (e *employeeRepository) FindByPhone(phone string) (model.Employee, error) {
	r := e.db.QueryRow("SELECT id, name, phone_number, address FROM employee WHERE phone_number=$1;", phone)

	emp := model.Employee{}
	err := r.Scan(&emp.ID, &emp.Name, &emp.PhoneNumber, &emp.Address)
	if err != nil {
		return emp, err
	}

	return emp, nil
}

// DeleteById implements EmployeeRepository.
func (e *employeeRepository) DeleteById(id string) error {
	_, err := e.db.Exec("DELETE FROM employee WHERE id=$1;", id)
	if err != nil {
		return err
	}

	return nil
}

// FindByAll implements EmployeeRepository.
func (e *employeeRepository) FindByAll() ([]model.Employee, error) {
	r, err := e.db.Query("SELECT id, name, phone_number, address FROM employee;")
	if err != nil {
		return nil, err
	}

	defer r.Close()

	employees := []model.Employee{}
	for r.Next() {
		employee := model.Employee{}
		err = r.Scan(&employee.ID, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

// FindById implements EmployeeRepository.
func (e *employeeRepository) FindById(id string) (model.Employee, error) {
	r := e.db.QueryRow("SELECT id, name, phone_number, address FROM employee WHERE id=$1;", id)

	emp := model.Employee{}
	err := r.Scan(&emp.ID, &emp.Name, &emp.PhoneNumber, &emp.Address)
	if err != nil {
		return emp, err
	}

	return emp, nil
}

// Save implements EmployeeRepository.
func (e *employeeRepository) Save(emp model.Employee) error {
	_, err := e.db.Exec("INSERT INTO employee VALUES ($1, $2, $3, $4);", emp.ID, emp.Name, emp.PhoneNumber, emp.Address)
	if err != nil {
		return err
	}

	return nil
}

// UpdateById implements EmployeeRepository.
func (e *employeeRepository) UpdateById(emp model.Employee) error {
	_, err := e.db.Exec("UPDATE employee SET name=$1, phone_number=$2, address=$3  WHERE id=$4;", emp.Name, emp.PhoneNumber, emp.Address, emp.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
