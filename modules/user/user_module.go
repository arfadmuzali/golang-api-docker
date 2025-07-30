package user

import (
	"learn/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserModule(rg *gin.RouterGroup) {
	user := rg.Group("/user", middleware.JWTAuthMiddleware())

	UserController(user)
}
