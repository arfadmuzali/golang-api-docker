package modules

import (
	"github.com/gin-gonic/gin"
	"learn/modules/auth"
	"learn/modules/hello"
)

func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("/")
	auth.RegisterAuthModules(api)
	hello.RegisterHelloModule(api)
}
