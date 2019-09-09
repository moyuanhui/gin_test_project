package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Book struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:check_out binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// v9自定义验证

// v9 多语言参数验证

// func bookDateEnable(v *validator.Validate) bool {

// 	if date,ok := field

// 	return true
// }

func main() {
	r := gin.Default()
	r.GET("/booking", func(ctx *gin.Context) {
		var b Book
		if err := ctx.ShouldBind(&b); err != nil {
			ctx.String(http.StatusInternalServerError, "Error:%v", err)
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":      http.StatusOK,
			"message":   "OK",
			"meta_data": fmt.Sprintf("%v", b),
		})

	})
	r.Run()
}
