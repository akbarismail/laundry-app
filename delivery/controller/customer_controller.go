package controller

import (
	"clean-code/model"
	"clean-code/usecase"
	"fmt"
	"os"
)

type CustomerController struct {
	customerUseCase usecase.CustomerUseCase
}

func (c *CustomerController) CustomerMenuForm() {
	fmt.Println(`
	|		+++++ Master Customer +++++	|
	| 1. Add Data					|
	| 2. Show Data					|
	| 3. Update Data				|
	| 4. Delete Data				|
	| 5. Show By Phone Data			|
	| 6. Exit						|
	`)

	fmt.Print("Choose Menu (1-6) *don't press space keyboard: ")
	var selectMenuCustomer string
	fmt.Scanln(&selectMenuCustomer)

	switch selectMenuCustomer {
	case "1":
		c.insertFormCustomer()
	case "2":
		c.showListFormCustomer()
	case "3":
		c.updateFormCustomer()
	case "4":
		c.deleteFormCustomer()
	case "5":
		c.showPhoneFormCustomer()
	case "6":
		os.Exit(0)
	}
}

func (c *CustomerController) showPhoneFormCustomer() {
	var customer model.Customer
	fmt.Print("Input Phone Number: ")
	fmt.Scanln(&customer.PhoneNumber)

	c2, err := c.customerUseCase.GetByPhone(customer.PhoneNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("ID: %s, Name: %s, Phone Number: %s, Address: %s\n", c2.ID, c2.Name, c2.PhoneNumber, c2.Address)
}

func (c *CustomerController) deleteFormCustomer() {
	var customer model.Customer
	fmt.Print("Input ID: ")
	fmt.Scanln(&customer.ID)

	err := c.customerUseCase.Delete(customer.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *CustomerController) updateFormCustomer() {
	var customer model.Customer
	fmt.Print("Input ID: ")
	fmt.Scanln(&customer.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("Input Phone Number: ")
	fmt.Scanln(&customer.PhoneNumber)
	fmt.Print("Input Address: ")
	fmt.Scanln(&customer.Address)

	err := c.customerUseCase.Update(customer)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *CustomerController) showListFormCustomer() {
	customers, err := c.customerUseCase.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(customers) == 0 {
		fmt.Println("Customer is empty")
		return
	}

	for _, customer := range customers {
		fmt.Printf("ID: %s, Name: %s, Phone Number: %s, Address: %s\n", customer.ID, customer.Name, customer.PhoneNumber, customer.Address)
	}
}

func (c *CustomerController) insertFormCustomer() {
	var customer model.Customer
	fmt.Print("Input ID: ")
	fmt.Scanln(&customer.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("Input Phone Number: ")
	fmt.Scanln(&customer.PhoneNumber)
	fmt.Print("Input Address: ")
	fmt.Scanln(&customer.Address)

	err := c.customerUseCase.CreateNew(customer)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewCustomerController(customerUseCase usecase.CustomerUseCase) *CustomerController {
	return &CustomerController{
		customerUseCase: customerUseCase,
	}
}
