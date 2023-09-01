package repository

import (
	"clean-code/model"
	"database/sql"
)

type BillDetailRepository interface {
	Save(billDetail model.BillDetail) error
}

type billDetailRepository struct {
	db *sql.DB
}

// Save implements BillDetailRepository.
func (bd *billDetailRepository) Save(billDetail model.BillDetail) error {
	_, err := bd.db.Exec("INSERT INTO bill_detail VALUES($1, $2, $3, $4, $5);", billDetail.ID, billDetail.Bill.ID, billDetail.Product.ID, billDetail.ProductPrice, billDetail.Qty)
	if err != nil {
		return err
	}

	return nil
}

func NewBillDetailRepository(db *sql.DB) BillDetailRepository {
	return &billDetailRepository{db: db}
}
