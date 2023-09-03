package controller

import (
	"bufio"
	"clean-code/model"
	"clean-code/usecase"
	"fmt"
	"os"
	"time"
)

type BillController struct {
	billUseCase usecase.BillUseCase
}

var layout = "2006-01-02"
var amount int

func (b *BillController) BillMenuForm() {
	fmt.Println(`
	|		+++++ Master Bill +++++	|
	| 1. Add Data					|
	| 2. Show List Data				|
	| 3. Show by Id Data			|
	| 4. Exit						|
	`)

	fmt.Print("Choose Menu (1-4) *don't press space keyboard: ")
	var selectMenuBill string
	fmt.Scanln(&selectMenuBill)

	switch selectMenuBill {
	case "1":
		b.insertFormProduct()
	case "2":
		b.showListFormBills()
	case "3":
		b.showIdFormBill()
	case "4":
		os.Exit(0)
	}
}

func (b *BillController) showIdFormBill() {
	var billModel model.Bill
	fmt.Print("Input ID: ")
	fmt.Scanln(&billModel.ID)

	bill, billsDetail, err := b.billUseCase.GetById(billModel.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	if bill.ID != billModel.ID {
		fmt.Println("Bill ID is Not Found")
		return
	}

	fmt.Println("===================================")
	fmt.Printf("ID: %s,\nBill Date: %s,\nEmployee Name: %s,\nEmployee Phone Number: %s,\nCustomer ID: %s,\nCustomer Name: %s,\nCustomer Phone Number: %s,\nCustomer Address: %s\n", bill.ID, bill.BillDate.Format(layout), bill.Employee.Name, bill.Employee.PhoneNumber, bill.Customer.ID, bill.Customer.Name, bill.Customer.PhoneNumber, bill.Customer.Address)

	for _, bd := range billsDetail {
		amount = bd.ProductPrice * bd.Qty
		fmt.Println("===================================")
		fmt.Printf("Bill Detail ID: %s,\nBill ID: %s,\nEntry Date: %s,\nFinish Date: %s,\nProduct Name: %s,\nUnit: %s,\nProduct Price: %d,\nQTY: %d,\nGrand Total: %d\n", bd.ID, bd.Bill.ID, bd.Bill.EntryDate.Format(layout), bd.Bill.FinishDate.Format(layout), bd.Product.Name, bd.Product.Uom.Name, bd.ProductPrice, bd.Qty, amount)
	}
}

func (b *BillController) showListFormBills() {
	bills, billsDetail, err := b.billUseCase.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(bills) == 0 {
		fmt.Println("Transaction is empty")
		return
	}

	for _, bill := range bills {
		fmt.Println("===================================")
		fmt.Printf("ID: %s,\nBill Date: %s,\nEmployee Name: %s,\nEmployee Phone Number: %s,\nCustomer ID: %s,\nCustomer Name: %s,\nCustomer Phone Number: %s,\nCustomer Address: %s\n", bill.ID, bill.BillDate.Format(layout), bill.Employee.Name, bill.Employee.PhoneNumber, bill.Customer.ID, bill.Customer.Name, bill.Customer.PhoneNumber, bill.Customer.Address)

	}

	for _, bd := range billsDetail {
		amount = bd.ProductPrice * bd.Qty
		fmt.Println("===================================")
		fmt.Printf("Bill Detail ID: %s,\nBill ID: %s,\nEntry Date: %s,\nFinish Date: %s,\nProduct Name: %s,\nUnit: %s,\nProduct Price: %d,\nQTY: %d,\nGrand Total: %d\n", bd.ID, bd.Bill.ID, bd.Bill.EntryDate.Format(layout), bd.Bill.FinishDate.Format(layout), bd.Product.Name, bd.Product.Uom.Name, bd.ProductPrice, bd.Qty, amount)
	}
}

func (b *BillController) insertFormProduct() {
	var bill model.Bill
	var billDetail model.BillDetail
	myScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input ID: ")
	fmt.Scanln(&bill.ID)

	// fmt.Print("Input Bill Date (YYY-MM-DD): ")
	bill.BillDate = time.Now()

	// fmt.Print("Input Entry Date (YYY-MM-DD): ")
	bill.EntryDate = time.Now()

	fmt.Print("Input Finish Date (YYY-MM-DD): ")
	myScanner.Scan()
	bill.FinishDate, _ = time.Parse(layout, myScanner.Text())

	fmt.Print("Input Employee ID: ")
	fmt.Scanln(&bill.Employee.ID)
	fmt.Print("Input Customer ID: ")
	fmt.Scanln(&bill.Customer.ID)

	fmt.Print("Input Bill Detail ID: ")
	fmt.Scanln(&billDetail.ID)
	fmt.Print("Input Bill ID: ")
	fmt.Scanln(&billDetail.Bill.ID)
	fmt.Print("Input Product ID: ")
	fmt.Scanln(&billDetail.Product.ID)
	fmt.Print("Input Product Price: ")
	fmt.Scanln(&billDetail.ProductPrice)
	fmt.Print("Input QTY: ")
	fmt.Scanln(&billDetail.Qty)

	err := b.billUseCase.CreateNew(bill, billDetail)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewBillController(billUseCase usecase.BillUseCase) *BillController {
	return &BillController{
		billUseCase: billUseCase,
	}
}
