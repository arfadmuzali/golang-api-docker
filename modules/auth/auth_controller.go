package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthController(api *gin.RouterGroup) {

	api.POST("/login", func(ctx *gin.Context) {
		var body LoginDto
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}

		result, err := HandleLogin(&body)
		if err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}
		ctx.JSON(200, APIResponse{Status: 200, Message: "success", Data: result})
	})

	api.POST("/register", func(ctx *gin.Context) {
		var body RegisterDto
		fmt.Println("register result")
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}

		result, err := HandleRegister(&body)
		if err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}
		ctx.JSON(200, APIResponse{Status: 200, Message: "success", Data: result})
	})
}
