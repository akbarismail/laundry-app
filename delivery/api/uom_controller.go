package api

import (
	"clean-code/model"
	"clean-code/usecase"
	"clean-code/util/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UomController struct {
	uomUseCase  usecase.UomUseCase
	routerGroup *gin.RouterGroup
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

	uom.ID = common.GenerateID()
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

func (u *UomController) Route() {
	u.routerGroup.PATCH("/uoms/:id", u.updateUom)
	u.routerGroup.DELETE("/uoms/:id", u.deleteUom)
	u.routerGroup.GET("/uoms", u.findUoms)
	u.routerGroup.POST("/uoms", u.createUom)

}

func NewUomController(uomUseCase usecase.UomUseCase, routerGroup *gin.RouterGroup) *UomController {
	return &UomController{
		uomUseCase:  uomUseCase,
		routerGroup: routerGroup,
	}
}
