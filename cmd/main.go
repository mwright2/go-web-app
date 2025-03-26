package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	templates "mjuniorw/web/templates/pages"
)

func render(status int, c *gin.Context, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		render(http.StatusOK, c, templates.Home())
	})

	router.Run(":8080")
}
