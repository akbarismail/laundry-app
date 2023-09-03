package controller

import (
	"clean-code/model"
	"clean-code/usecase"
	"fmt"
	"os"
)

type UomController struct {
	uomUseCase usecase.UomUseCase
}

func (u *UomController) UomMenuForm() {
	fmt.Println(`
	|		+++++ Master UOM +++++	|
	| 1. Add Data			   	    |
	| 2. Show Data					|
	| 3. Update Data				|
	| 4. Delete Data				|
	| 5. Show By Name Data      	|
	| 6. Exit                       |
	`)

	fmt.Print("Choose Menu (1-6) *don't press space keyboard: ")
	var selectMenuUom string
	fmt.Scanln(&selectMenuUom)

	switch selectMenuUom {
	case "1":
		u.insertFormUom()
	case "2":
		u.showListFormUom()
	case "3":
		u.updateFormUom()
	case "4":
		u.deleteFormUom()
	case "5":
		u.showNameFormUom()
	case "6":
		os.Exit(0)
	}

}

func (u *UomController) insertFormUom() {
	var uom model.Uom
	fmt.Print("Input ID: ")
	fmt.Scanln(&uom.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&uom.Name)

	err := u.uomUseCase.CreateNew(uom)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (u *UomController) showListFormUom() {
	uoms, err := u.uomUseCase.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(uoms) == 0 {
		fmt.Println("Uom is empty")
		return
	}

	for _, uom := range uoms {
		fmt.Printf("ID: %s, Name: %s\n", uom.ID, uom.Name)
	}
}

func (u *UomController) updateFormUom() {
	var uom model.Uom
	fmt.Print("Input ID: ")
	fmt.Scanln(&uom.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&uom.Name)

	err := u.uomUseCase.Update(uom)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (u *UomController) deleteFormUom() {
	var uom model.Uom
	fmt.Print("Input ID: ")
	fmt.Scanln(&uom.ID)

	err := u.uomUseCase.Delete(uom.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (u *UomController) showNameFormUom() {
	var uom model.Uom
	fmt.Print("Input Name: ")
	fmt.Scanln(&uom.Name)

	uoms, err := u.uomUseCase.GetByName(uom.Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(uoms) == 0 {
		fmt.Println("Uom is empty")
		return
	}

	for _, uomValue := range uoms {
		fmt.Printf("ID: %s, Name: %s\n", uomValue.ID, uomValue.Name)
	}
}

func NewUomController(uomUseCase usecase.UomUseCase) *UomController {
	return &UomController{
		uomUseCase: uomUseCase,
	}
}
