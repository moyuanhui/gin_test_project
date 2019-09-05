package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(ctx *gin.Context) {
		ctx.String(200, "GET")
	})

	r.POST("/post", func(ctx *gin.Context) {
		ctx.String(200, "POST")
	})
	r.Run(":5001")
}
