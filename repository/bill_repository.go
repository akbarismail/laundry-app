package repository

import (
	"clean-code/model"
	"database/sql"
)

type BillRepository interface {
	Save(bill model.Bill) error
	FindByAll() ([]model.Bill, error)
	FindById(id string) (model.Bill, error)
}

type billRepository struct {
	db *sql.DB
}

// FindById implements BillRepository.
func (b *billRepository) FindById(id string) (model.Bill, error) {
	r := b.db.QueryRow("SELECT b.id, b.bill_date, b.entry_date, b.finish_date, e.id, e.name, e.phone_number, e.address, c.id, c.name, c.phone_number, c.address FROM bill AS b JOIN employee AS e ON e.id = b.employee_id JOIN customer AS c ON c.id = b.customer_id WHERE b.id=$1;", id)

	bill := model.Bill{}
	err := r.Scan(&bill.ID, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.Employee.ID, &bill.Employee.Name, &bill.Employee.PhoneNumber, &bill.Employee.Address, &bill.Customer.ID, &bill.Customer.Name, &bill.Customer.PhoneNumber, &bill.Customer.Address)
	if err != nil {
		return bill, err
	}
	return bill, nil
}

// FindByAll implements BillRepository.
func (b *billRepository) FindByAll() ([]model.Bill, error) {
	r, err := b.db.Query("SELECT b.id, b.bill_date, b.entry_date, b.finish_date, e.id, e.name, e.phone_number, e.address, c.id, c.name, c.phone_number, c.address FROM bill AS b JOIN employee AS e ON e.id = b.employee_id JOIN customer AS c ON c.id = b.customer_id;")
	if err != nil {
		return nil, err
	}

	defer r.Close()
	bills := []model.Bill{}

	for r.Next() {
		bill := model.Bill{}
		err = r.Scan(&bill.ID, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.Employee.ID, &bill.Employee.Name, &bill.Employee.PhoneNumber, &bill.Employee.Address, &bill.Customer.ID, &bill.Customer.Name, &bill.Customer.PhoneNumber, &bill.Customer.Address)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}

	return bills, nil
}

// Save implements BillRepository.
func (b *billRepository) Save(bill model.Bill) error {
	_, err := b.db.Exec("INSERT INTO bill VALUES($1, $2, $3, $4, $5, $6);", bill.ID, bill.BillDate, bill.EntryDate, bill.FinishDate, bill.Employee.ID, bill.Customer.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{db: db}
}
