package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})
	r.Run()
}
