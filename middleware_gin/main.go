package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultErrorWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())
	r.GET("/test", func(ctx *gin.Context) {
		// ctx.String(http.StatusOK, "ok")
		name := ctx.DefaultQuery("name", "lal")
		panic("test panic")
		ctx.String(http.StatusOK, "%s", name)
	})
	r.Run(":8080")
}
