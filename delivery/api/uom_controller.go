package api

import (
	"clean-code/delivery/middleware"
	"clean-code/model"
	"clean-code/model/dto"
	"clean-code/usecase"
	"clean-code/util/common"
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))

	rows, paging, err := u.uomUseCase.Paging(dto.PageRequest{
		Page: page,
		Size: size,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get uoms",
		"data":    rows,
		"paging":  paging,
	})
}

func (u *UomController) Route() {
	u.routerGroup.PATCH("/uoms/:id", middleware.AuthMiddleWare(), u.updateUom)
	u.routerGroup.DELETE("/uoms/:id", middleware.AuthMiddleWare(), u.deleteUom)
	u.routerGroup.GET("/uoms", middleware.AuthMiddleWare(), u.findUoms)
	u.routerGroup.POST("/uoms", middleware.AuthMiddleWare(), u.createUom)

}

func NewUomController(uomUseCase usecase.UomUseCase, routerGroup *gin.RouterGroup) *UomController {
	return &UomController{
		uomUseCase:  uomUseCase,
		routerGroup: routerGroup,
	}
}
