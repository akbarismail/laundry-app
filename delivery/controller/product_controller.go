package controller

import (
	"clean-code/model"
	"clean-code/usecase"
	"fmt"
	"os"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func (p *ProductController) ProductMenuForm() {
	fmt.Println(`
	|		+++++ Master Product +++++	|
	| 1. Add Data					|
	| 2. Show Data					|
	| 3. Update Data				|
	| 4. Delete Data				|
	| 5. Show By Name Data			|
	| 6. Exit						|
	`)

	fmt.Print("Choose Menu (1-6) *don't press space keyboard: ")
	var selectMenuProduct string
	fmt.Scanln(&selectMenuProduct)

	switch selectMenuProduct {
	case "1":
		p.insertFormProduct()
	case "2":
		p.showListFormProduct()
	case "3":
		p.updateFormProduct()
	case "4":
		p.deleteFormProduct()
	case "5":
		p.showNameFormProduct()
	case "6":
		os.Exit(0)
	}
}

func (p *ProductController) insertFormProduct() {
	var product model.Product
	fmt.Print("Input ID: ")
	fmt.Scanln(&product.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&product.Name)
	fmt.Print("Input Price: ")
	fmt.Scanln(&product.Price)
	fmt.Print("Input Uom ID: ")
	fmt.Scanln(&product.Uom.ID)

	err := p.productUseCase.CreateNew(product)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (p *ProductController) showListFormProduct() {
	p2, err := p.productUseCase.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(p2) == 0 {
		fmt.Println("Product is empty")
		return
	}

	for _, productVal := range p2 {
		fmt.Printf("ID: %s, Name: %s, Price: %d, Uom ID: %s, Uom Name: %s\n", productVal.ID, productVal.Name, productVal.Price, productVal.Uom.ID, productVal.Uom.Name)
	}
}

func (p *ProductController) updateFormProduct() {
	var product model.Product
	fmt.Print("Input ID: ")
	fmt.Scanln(&product.ID)
	fmt.Print("Input Name: ")
	fmt.Scanln(&product.Name)
	fmt.Print("Input Price: ")
	fmt.Scanln(&product.Price)
	fmt.Print("Input Uom ID: ")
	fmt.Scanln(&product.Uom.ID)

	err := p.productUseCase.Update(product)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (p *ProductController) deleteFormProduct() {
	var product model.Product
	fmt.Print("Input ID: ")
	fmt.Scanln(&product.ID)

	err := p.productUseCase.Delete(product.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (p *ProductController) showNameFormProduct() {
	var product model.Product
	fmt.Print("Input Name: ")
	fmt.Scanln(&product.Name)

	p2, err := p.productUseCase.GetByName(product.Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(p2) == 0 {
		fmt.Println("Product is empty")
		return
	}

	for _, productVal := range p2 {
		fmt.Printf("ID: %s, Name: %s, Price: %d, Uom ID: %s, Uom Name: %s\n", productVal.ID, productVal.Name, productVal.Price, productVal.Uom.ID, productVal.Uom.Name)
	}
}

func NewProductController(productUseCase usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: productUseCase,
	}
}
