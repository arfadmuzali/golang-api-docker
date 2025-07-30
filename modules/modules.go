package modules

import (
	"learn/modules/auth"
	"learn/modules/hello"
	"learn/modules/user"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("/")
	auth.RegisterAuthModules(api)
	hello.RegisterHelloModule(api)
	user.RegisterUserModule(api)
}
