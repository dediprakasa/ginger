package main

import (
	"github.com/dediprakasa/ginger/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userRouter := r.Group("/user")

	{
		userController := new(controllers.UserController)
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
