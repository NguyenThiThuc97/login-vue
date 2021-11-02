package handler

import (
	"net/http"
	controllers "workspace/login/api/controllers"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	// h := &Handler{}

	g := c.R.Group("/api")

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	g.GET("/user/:id", controllers.FindUser)
	g.POST("/register", controllers.Register)
	g.POST("/signin", controllers.SignIn)

	productGroup := g.Group("/product")
	productGroup.POST("/create", controllers.CreateProduct)
	productGroup.DELETE("/delete/:id", controllers.DeleteProduct)

	c.R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Page Not Found",
		})
	})
}
