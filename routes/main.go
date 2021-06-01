package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	SetupUserRouter(r)
}
