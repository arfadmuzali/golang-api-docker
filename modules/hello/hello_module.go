package hello

import (
	"learn/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterHelloModule(rg *gin.RouterGroup) {
	hello := rg.Group("/hello", middleware.JWTAuthMiddleware())

	HelloController(hello)
}
