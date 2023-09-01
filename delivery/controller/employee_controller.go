package controller

import (
	"clean-code/model"
	"clean-code/usecase"
	"fmt"
	"os"
)

type EmployeeController struct {
	employeeUseCase usecase.EmployeeUseCase
}

func (e *EmployeeController) EmployeeMenuForm() {
	fmt.Println(`
	|		+++++ Master Employee +++++	|
	| 1. Add Data					|
	| 2. Show Data					|
	| 3. Update Data				|
	| 4. Delete Data				|
	| 5. Show By Phone Data			|
	| 6. Exit						|
	`)

	fmt.Print("Choose Menu (1-6): ")
	var selectMenuUom string
	fmt.Scanln(&selectMenuUom)

	switch selectMenuUom {
	case "1":
		e.insertFormEmployee()
	case "2":
		e.showListFormEmployee()
	case "3":
		e.updateFormEmployee()
	case "4":
		e.deleteFormEmployee()
	case "5":
		e.showPhoneFormEmployee()
	case "6":
		os.Exit(0)
	}
}

func (e *EmployeeController) showPhoneFormEmployee() {
	var employee model.Employee
	fmt.Print("Input Phone Number: ")
	fmt.Scanln(&employee.PhoneNumber)

	e2, err := e.employeeUseCase.GetByPhone(employee.PhoneNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("ID: %s, Name: %s, Phone Number: %s, Address: %s\n", e2.ID, e2.Name, e2.PhoneNumber, e2.Address)
}

func (e *EmployeeController) deleteFormEmployee() {
	var employee model.Employee
	fmt.Print("Input ID: ")
	fmt.Scanln(&employee.ID)

	err := e.employeeUseCase.Delete(employee.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (e *EmployeeController) updateFormEmployee() {
	var employee model.Employee
	fmt.Print("Input ID: ")
	fmt.Scanln(&employee.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&employee.Name)
	fmt.Print("Input Phone Number: ")
	fmt.Scanln(&employee.PhoneNumber)
	fmt.Print("Input Address: ")
	fmt.Scanln(&employee.Address)

	err := e.employeeUseCase.Update(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (e *EmployeeController) showListFormEmployee() {
	employees, err := e.employeeUseCase.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(employees) == 0 {
		fmt.Println("Customer is empty")
		return
	}

	for _, employee := range employees {
		fmt.Printf("ID: %s, Name: %s, Phone Number: %s, Address: %s\n", employee.ID, employee.Name, employee.PhoneNumber, employee.Address)
	}
}

func (e *EmployeeController) insertFormEmployee() {
	var employee model.Employee
	fmt.Print("Input ID: ")
	fmt.Scanln(&employee.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&employee.Name)
	fmt.Print("Input Phone Number: ")
	fmt.Scanln(&employee.PhoneNumber)
	fmt.Print("Input Address: ")
	fmt.Scanln(&employee.Address)

	err := e.employeeUseCase.CreateNew(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewEmployeeController(employeeUseCase usecase.EmployeeUseCase) *EmployeeController {
	return &EmployeeController{
		employeeUseCase: employeeUseCase,
	}
}
