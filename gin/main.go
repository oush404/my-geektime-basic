package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.Use(func(context *gin.Context) {
		println("first middleware")
	}, func(context *gin.Context) {
		println("second middleware")
	})
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world")
	})
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello, "+name)
	})
	//http://localhost:8080/order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "订单 ID 是 "+id)
	})
	server.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html")
		ctx.String(http.StatusOK, "view 是 "+view)
	})

	//POST
	server.POST("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, login")
	})

	server.Run(":8080")
}
