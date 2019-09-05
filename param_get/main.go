package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("test", func(ctx *gin.Context) {
		var firstName = ctx.Query("firstName")
		var lastName = ctx.DefaultQuery("lastName", "lastNamea")
		ctx.JSON(200, gin.H{
			"firstName": firstName,
			"lastName":  lastName,
		})
	})
	r.Run()
}
