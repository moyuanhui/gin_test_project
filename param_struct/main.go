package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2016-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
}

func testing(ctx *gin.Context) {

	var person Person
	if ctx.ShouldBind(&person) == nil {
		ctx.String(http.StatusOK, "%v", person)
	} else {
		ctx.String(http.StatusOK, "%v", person)
	}
}
