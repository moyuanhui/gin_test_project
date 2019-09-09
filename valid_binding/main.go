package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Age     string `form:"age" binding:"required,gt=10"`
	Address string `form:"address" binding:"required"`
	Name    string `form:"name" binding:"required"`
}

func main() {

	r := gin.Default()
	r.GET("/valid", func(ctx *gin.Context) {
		var person Person
		if err := ctx.ShouldBind(&person); err != nil {
			ctx.String(http.StatusInternalServerError, "Error:%v", err)
			ctx.Abort()
			return
		}
		ctx.String(http.StatusOK, "%v", person)
	})
	r.Run()
}
