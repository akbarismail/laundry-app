package api

import (
	"clean-code/model"
	"clean-code/usecase"
	"clean-code/util/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeUseCase usecase.EmployeeUseCase
	routerGroup     *gin.RouterGroup
}

func (e *EmployeeController) deleteEmployee(c *gin.Context) {
	id := c.Param("id")

	if err := e.employeeUseCase.Delete(id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete employee id: " + id,
	})
}

func (e *EmployeeController) updateEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := c.Param("id")
	employee.ID = id

	if err := e.employeeUseCase.Update(employee); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated employee id: " + id,
		"data":    employee,
	})
}

func (e *EmployeeController) findByPhoneEmployee(c *gin.Context) {
	phone := c.Param("phone")

	employee, err := e.employeeUseCase.GetByPhone(phone)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully",
		"data":    employee,
	})
}

func (e *EmployeeController) findEmployees(c *gin.Context) {
	employees, err := e.employeeUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully",
		"data":    employees,
	})
}

func (e *EmployeeController) createEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	employee.ID = common.GenerateID()
	if err := e.employeeUseCase.CreateNew(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created employee",
		"data":    employee,
	})
}

func (e *EmployeeController) Route() {
	e.routerGroup.DELETE("/employees/:id", e.deleteEmployee)
	e.routerGroup.PATCH("/employees/:id", e.updateEmployee)
	e.routerGroup.GET("/employees", e.findEmployees)
	e.routerGroup.GET("/employees/:phone", e.findByPhoneEmployee)
	e.routerGroup.POST("/employees", e.createEmployee)

}

func NewEmployeeController(employeeUseCase usecase.EmployeeUseCase, routerGroup *gin.RouterGroup) *EmployeeController {
	return &EmployeeController{
		employeeUseCase: employeeUseCase,
		routerGroup:     routerGroup,
	}
}
