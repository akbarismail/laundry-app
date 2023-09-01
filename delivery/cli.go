package delivery

import (
	"clean-code/config"
	"clean-code/delivery/controller"
	"clean-code/repository"
	"clean-code/usecase"
	"fmt"
	"os"
)

type Console struct {
	uomUseCase      usecase.UomUseCase
	productUseCase  usecase.ProductUseCase
	employeeUseCase usecase.EmployeeUseCase
	customerUseCase usecase.CustomerUseCase
	billUseCase     usecase.BillUseCase
}

func (c *Console) showMainMenu() {
	fmt.Println(`
	|+++++ Enigma Laundry Menu +++++|
	| 1. Master UOM                 |
	| 2. Master Product             |
	| 3. Master Customer            |
	| 4. Master Employee            |
	| 5. Transaction                |
	| 6. Exit                       |
	`)
	fmt.Print("Choose Menu (1-6): ")
}

func (c *Console) Run() {
	for {
		c.showMainMenu()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)

		switch selectedMenu {
		case "1":
			controller.NewUomController(c.uomUseCase).UomMenuForm()
		case "2":
			controller.NewProductController(c.productUseCase).ProductMenuForm()
		case "3":
			controller.NewCustomerController(c.customerUseCase).CustomerMenuForm()
		case "4":
			controller.NewEmployeeController(c.employeeUseCase).EmployeeMenuForm()
		case "5":
			controller.NewBillController(c.billUseCase)
		case "6":
			os.Exit(0)
		}
	}
}

func NewConsole() *Console {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	con, err := config.NewDBConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}

	db := con.Conn()

	uomRepository := repository.NewUomRepository(db)
	productRepository := repository.NewProductRepository(db)
	customerRepository := repository.NewCustomerRepository(db)
	employeeRepository := repository.NewEmployeeRepository(db)
	billRepository := repository.NewBillRepository(db)

	uomUseCase := usecase.NewUomUseCase(uomRepository)
	productUseCase := usecase.NewProductUseCase(productRepository, uomUseCase)
	customerUseCase := usecase.NewCustomerUseCase(customerRepository)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository)
	billUseCase := usecase.NewBillUseCase(billRepository, employeeUseCase, customerUseCase, productUseCase)

	return &Console{
		uomUseCase:      uomUseCase,
		productUseCase:  productUseCase,
		employeeUseCase: employeeUseCase,
		customerUseCase: customerUseCase,
		billUseCase:     billUseCase,
	}
}
