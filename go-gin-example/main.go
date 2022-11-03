package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	})
	g.GET("/user/info", func(c *gin.Context) {
		user := c.Query("user")
		name := c.Query("name")
		c.JSON(200, gin.H{
			"艾迪": user,
			"名字": name,
			"11": "23",
		})
	})
	g.GET("/user/info/:user/:name", func(c *gin.Context) {
		user := c.Param("user")
		name := c.Param("name")
		c.JSON(200, gin.H{
			"艾迪": user,
			"名字": name,
		})
	})
	g.LoadHTMLGlob("templates/**/*")
	g.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	g.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "users/index",
		})
	})
	g.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	g.LoadHTMLFiles("templates/index.tmpl")
	g.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://wdx'>wodelog</a>")
	})

	err := g.Run(":8080")
	if err != nil {
		return
	}
}
