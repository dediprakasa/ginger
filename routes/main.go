package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	SetupUserRouter(r)
	return r
}
