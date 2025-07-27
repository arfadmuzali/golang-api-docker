package main

import (
	"fmt"
	"learn/modules"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Failed load env file")
		return
	}
	r := gin.Default()

	modules.RegisterAPIRoutes(r)
	r.Run(":3000")
}
