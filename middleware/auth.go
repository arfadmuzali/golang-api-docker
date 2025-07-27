package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		secretKey := os.Getenv("AUTH_SECRET")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized", "status": 401})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"message": err.Error(), "status": 401})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(498, gin.H{"message": "Invalid auth token", "status": 498})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("AuthUser", claims)
		} else {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized", "status": 401})
			return
		}
		ctx.Next()
	}

}
