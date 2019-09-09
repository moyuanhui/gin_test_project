package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPWhiteHandleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		iplist := []string{
			"127.0.0.1",
		}
		clientIp := c.ClientIP()
		flag := false
		for _, host := range iplist {
			if host == clientIp {
				flag = true
				break
			}
		}

		if !flag {
			c.String(http.StatusOK, "%s,%s", "ip not in iplist,clientIP: ", clientIp)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(IPWhiteHandleFunc())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
		})
	})
	r.Run()
}
