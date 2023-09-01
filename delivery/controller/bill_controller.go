package controller

import "clean-code/usecase"

type BillController struct {
	billUseCase usecase.BillUseCase
}

func NewBillController(billUseCase usecase.BillUseCase) *BillController {
	return &BillController{
		billUseCase: billUseCase,
	}
}
