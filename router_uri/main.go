package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name": ctx.Param("name"),
			"id":   ctx.Param("id"),
		})
	})
	r.Run()
}
