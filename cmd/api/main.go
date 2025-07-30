package main

import (
	"learn/modules"
	"learn/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Failed load env file")
		return
	}

	err = utils.DBinit()
	if err != nil {
		log.Fatal("Error: Failed to connect DB")
		return
	}

	r := gin.Default()

	modules.RegisterAPIRoutes(r)
	r.Run(":3000")
}
