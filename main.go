package main

import "clean-code/delivery"

func main() {
	// delivery.NewConsole().Run()
	delivery.NewServer().Run()

	// middleware harus di panggil sebelum handler

	// r := gin.Default()

	// r.Use(func(ctx *gin.Context) {
	// 	fmt.Println("sebelum hello world")

	// 	ctx.Next()

	// 	fmt.Println("sesudah hello world")
	// })

	// r.Use(func(ctx *gin.Context) {
	// 	fmt.Println("sebelum hello world 2")

	// 	ctx.Next()

	// 	fmt.Println("sesudah hello world 2")
	// })

	// r.GET("/", func(ctx *gin.Context) {
	// 	fmt.Println("Hello World")
	// })

	// r.Run()
}
