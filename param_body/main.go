package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(ctx *gin.Context) {

		bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			ctx.String(500, err.Error())
			ctx.Abort()

		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := ctx.PostForm("firstName")
		lastName := ctx.DefaultPostForm("lastName", "lastNamedd")

		//ctx.String(http.StatusOK, string(bodyBytes))

		ctx.String(http.StatusOK, "%s %s %s", firstName, lastName, string(bodyBytes))
	})
	r.Run()
}
