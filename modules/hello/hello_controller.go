package hello

import "github.com/gin-gonic/gin"

func HelloController(api *gin.RouterGroup) {
	api.GET("/", func(ctx *gin.Context) {

		session, isExists := ctx.Get("AuthUser")
		if !isExists {
			ctx.JSON(500, gin.H{"message": "User data not found"})
			return
		}
		ctx.JSON(200, gin.H{"message": "this is protected route", "user_data": session})
	})

}
