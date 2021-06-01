package routes

import (
	"github.com/dediprakasa/ginger/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine) *gin.Engine {
	userRouter := r.Group("/user")
	{
		userController := new(controllers.UserController)
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}
	return r
}
