package repository

import (
	"clean-code/model"
	"database/sql"
)

type BillRepository interface {
	Save(bill model.Bill, billDetail model.BillDetail) error
	FindByAll() ([]model.Bill, []model.BillDetail, error)
	FindById(id string) (model.Bill, []model.BillDetail, error)
}

type billRepository struct {
	db *sql.DB
}

// FindById implements BillRepository.
func (b *billRepository) FindById(id string) (model.Bill, []model.BillDetail, error) {
	tx, err := b.db.Begin()
	if err != nil {
		return model.Bill{}, nil, err
	}

	r := tx.QueryRow("SELECT b.id, b.bill_date, b.entry_date, b.finish_date, e.id, e.name, e.phone_number, e.address, c.id, c.name, c.phone_number, c.address FROM bill AS b JOIN employee AS e ON e.id = b.employee_id JOIN customer AS c ON c.id = b.customer_id WHERE b.id=$1;", id)

	bill := model.Bill{}
	err = r.Scan(&bill.ID, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.Employee.ID, &bill.Employee.Name, &bill.Employee.PhoneNumber, &bill.Employee.Address, &bill.Customer.ID, &bill.Customer.Name, &bill.Customer.PhoneNumber, &bill.Customer.Address)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return bill, nil, err
		}

		return bill, nil, err
	}

	r2, err := tx.Query("SELECT bd.id, bd.product_price, bd.qty, b.id, b.bill_date, b.entry_date, b.finish_date, p.id, p.name, p.price, u.id, u.name FROM bill_detail AS bd JOIN bill AS b ON b.id=bd.bill_id JOIN product AS p ON p.id=bd.product_id JOIN uom as u ON u.id=p.uom_id WHERE bd.bill_id=$1;", id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return bill, nil, err
		}

		return bill, nil, err
	}

	defer r2.Close()
	billsDetail := []model.BillDetail{}
	billDetail := model.BillDetail{}
	for r2.Next() {
		err = r2.Scan(&billDetail.ID, &billDetail.ProductPrice, &billDetail.Qty, &billDetail.Bill.ID, &billDetail.Bill.BillDate, &billDetail.Bill.EntryDate, &billDetail.Bill.FinishDate, &billDetail.Product.ID, &billDetail.Product.Name, &billDetail.Product.Price, &billDetail.Product.Uom.ID, &billDetail.Product.Uom.Name)
		if err != nil {
			return bill, nil, err
		}

		billsDetail = append(billsDetail, billDetail)
	}

	err = tx.Commit()
	if err != nil {
		return bill, nil, err
	}

	return bill, billsDetail, nil
}

// FindByAll implements BillRepository.
func (b *billRepository) FindByAll() ([]model.Bill, []model.BillDetail, error) {
	tx, err := b.db.Begin()
	if err != nil {
		return nil, nil, err
	}

	r, err := tx.Query("SELECT b.id, b.bill_date, b.entry_date, b.finish_date, e.id, e.name, e.phone_number, e.address, c.id, c.name, c.phone_number, c.address FROM bill AS b JOIN employee AS e ON e.id = b.employee_id JOIN customer AS c ON c.id = b.customer_id;")
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, nil, err
		}

		return nil, nil, err
	}

	defer r.Close()
	bills := []model.Bill{}
	bill := model.Bill{}

	for r.Next() {
		err = r.Scan(&bill.ID, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.Employee.ID, &bill.Employee.Name, &bill.Employee.PhoneNumber, &bill.Employee.Address, &bill.Customer.ID, &bill.Customer.Name, &bill.Customer.PhoneNumber, &bill.Customer.Address)
		if err != nil {
			return nil, nil, err
		}
		bills = append(bills, bill)
	}

	r2, err := tx.Query("SELECT bd.id, bd.product_price, bd.qty, b.id, b.bill_date, b.entry_date, b.finish_date, p.id, p.name, p.price, u.id, u.name FROM bill_detail AS bd JOIN bill AS b ON b.id=bd.bill_id JOIN product AS p ON p.id=bd.product_id JOIN uom as u ON u.id=p.uom_id;")
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, nil, err
		}

		return nil, nil, err
	}

	defer r2.Close()
	billsDetail := []model.BillDetail{}
	billDetail := model.BillDetail{}
	for r2.Next() {
		err = r2.Scan(&billDetail.ID, &billDetail.ProductPrice, &billDetail.Qty, &billDetail.Bill.ID, &billDetail.Bill.BillDate, &billDetail.Bill.EntryDate, &billDetail.Bill.FinishDate, &billDetail.Product.ID, &billDetail.Product.Name, &billDetail.Product.Price, &billDetail.Product.Uom.ID, &billDetail.Product.Uom.Name)
		if err != nil {
			return nil, nil, err
		}

		billsDetail = append(billsDetail, billDetail)
	}

	err = tx.Commit()
	if err != nil {
		return nil, nil, err
	}

	return bills, billsDetail, nil
}

// Save implements BillRepository.
func (b *billRepository) Save(bill model.Bill, billDetail model.BillDetail) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO bill VALUES($1, $2, $3, $4, $5, $6);", bill.ID, bill.BillDate, bill.EntryDate, bill.FinishDate, bill.Employee.ID, bill.Customer.ID)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}

		return err
	}

	_, err = tx.Exec("INSERT INTO bill_detail VALUES($1, $2, $3, $4, $5);", billDetail.ID, billDetail.Bill.ID, billDetail.Product.ID, billDetail.ProductPrice, billDetail.Qty)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{
		db: db,
	}
}
