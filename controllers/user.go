package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "register endpointtt",
	})
}

func (u UserController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login endpoint",
	})
}
