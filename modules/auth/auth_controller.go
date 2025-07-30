package auth

import (
	"github.com/gin-gonic/gin"
	"log"
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
		log.Println("result again")
		var body RegisterDto
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}

		result, err := HandleRegister(&body)
		log.Println("result again", err)
		if err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}
		ctx.JSON(200, APIResponse{Status: 200, Message: "success", Data: result})
	})
}
