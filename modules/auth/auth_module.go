package auth

import "github.com/gin-gonic/gin"

func RegisterAuthModules(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	AuthController(auth)
}
