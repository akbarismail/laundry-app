package controller

import (
	"clean-code/usecase"
	"fmt"
	"os"
)

type BillController struct {
	billUseCase usecase.BillUseCase
}

func (b *BillController) BillMenuForm() {
	fmt.Println(`
	|		+++++ Master Bill +++++	|
	| 1. Add Data					|
	| 2. Show List Data				|
	| 3. Show by Id Data			|
	| 4. Exit						|
	`)

	fmt.Print("Choose Menu (1-4): ")
	var selectMenuBill string
	fmt.Scanln(&selectMenuBill)

	switch selectMenuBill {
	case "1":

	case "2":

	case "3":

	case "4":

	case "5":

	case "6":
		os.Exit(0)
	}
}

func NewBillController(billUseCase usecase.BillUseCase) *BillController {
	return &BillController{
		billUseCase: billUseCase,
	}
}
