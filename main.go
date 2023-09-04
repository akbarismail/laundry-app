package main

import "clean-code/delivery"

func main() {
	// r := gin.Default()
	// api := r.Group("/api/v1")

	// api.GET("/status", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "successfully",
	// 	})
	// })

	// // get products by id
	// api.GET("/products/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "berhasil mendapatkan products dengan id: " + id,
	// 	})
	// })

	// // filter name products by query param
	// api.GET("/products", func(ctx *gin.Context) {
	// 	name := ctx.Query("name")

	// 	if name == "" {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "nama produk tidak di temukan",
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"message": "berhasil mencari products dengan nama: " + name,
	// 		})
	// 	}

	// })

	// api.GET("/uoms/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")

	// 	uom := model.Uom{
	// 		ID:   id,
	// 		Name: "KG",
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "successfully",
	// 		"data":    uom,
	// 	})
	// })

	// api.POST("/uoms", func(ctx *gin.Context) {
	// 	var uom model.Uom

	// 	if err := ctx.ShouldBindJSON(&uom); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, err.Error())
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusCreated, &uom)
	// })

	// r.Run()

	// delivery.NewConsole().Run()
	delivery.NewServer().Run()
}
