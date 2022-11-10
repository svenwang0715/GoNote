package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	})
	g.GET("/user/search/:name/:address", func(c *gin.Context) {
		username := c.Param("name")
		ad := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"地址":       ad,
		})
	})

	err := g.Run(":8080")
	if err != nil {
		return
	}
}
