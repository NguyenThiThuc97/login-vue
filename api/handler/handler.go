package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group("/api")

	// g.Use(ValidateToken())

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	g.POST("/signin", h.SignIn)

	c.R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Page Not Found",
		})
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	if c.Request.FormValue("username") == "admin" && c.Request.FormValue("password") == "123" {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"hello":  "it's login",
		})
	} else {
		c.JSON(401, gin.H{
			"status":  "err",
			"message": "incorrect info",
		})
	}

}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")

		if token == "" {
			c.JSON(401, gin.H{
				"message": "Token Required",
			})
			c.Abort()
			return
		}
		if token != "accesstoken" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
