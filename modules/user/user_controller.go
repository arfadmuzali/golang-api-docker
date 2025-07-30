package user

import "github.com/gin-gonic/gin"

func UserController(api *gin.RouterGroup) {

	api.GET("/", func(ctx *gin.Context) {
		result, err := GetAllUser()
		if err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}
		ctx.JSON(200, APIResponse{Status: 200, Message: "Success", Data: result})
	})

	api.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(400, APIResponse{Status: 400, Message: "Params id is required"})
			return
		}
		result, err := GetUser(id)
		if err != nil {
			ctx.JSON(400, APIResponse{Status: 400, Message: err.Error()})
			return
		}
		ctx.JSON(200, APIResponse{Status: 200, Message: "Success", Data: result})
	})

}
