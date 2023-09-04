package api

import (
	"clean-code/model"
	"clean-code/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UomController struct {
	uomUseCase usecase.UomUseCase
	engine     *gin.Engine
}

func (u *UomController) updateUom(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := c.Param("id")
	uom.ID = id

	err := u.uomUseCase.Update(uom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    uom,
		"message": "successfully updated uom id: " + id,
	})
}

func (u *UomController) deleteUom(c *gin.Context) {
	id := c.Param("id")

	err := u.uomUseCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete uom id: " + id,
	})
}

func (u *UomController) createUom(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	uom.ID = uuid.NewString()
	err := u.uomUseCase.CreateNew(uom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully",
	})
}

func (u *UomController) findUoms(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		u2, err := u.uomUseCase.GetAll()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": err.Error(),
				"data":    u2,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "successfully",
			"data":    u2,
		})
		return

	} else {
		u2, err := u.uomUseCase.GetByName(name)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": err.Error(),
				"data":    u2,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "successfully",
			"data":    u2,
		})

		return
	}
}

func NewUomController(uomUseCase usecase.UomUseCase, engine *gin.Engine) {
	uomController := UomController{
		uomUseCase: uomUseCase,
		engine:     engine,
	}

	rg := engine.Group("/api/v1")
	rg.PATCH("/uoms/:id", uomController.updateUom)
	rg.DELETE("/uoms/:id", uomController.deleteUom)
	rg.GET("/uoms", uomController.findUoms)
	rg.POST("/uoms", uomController.createUom)
}
