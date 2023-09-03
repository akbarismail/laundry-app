package model

import "time"

type Bill struct {
	ID         string
	BillDate   time.Time
	EntryDate  time.Time
	FinishDate time.Time
	Employee   Employee
	Customer   Customer
}

type BillDetail struct {
	ID           string
	Bill         Bill
	Product      Product
	ProductPrice int
	Qty          int
}
