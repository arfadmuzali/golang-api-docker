package main

import (
	"learn/utils"
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	redisOpt := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_ADDR")}

	if err != nil {
		log.Fatal("Error: Can't load .env file")
	}
	serv := asynq.NewServer(redisOpt, asynq.Config{Concurrency: 10})

	mux := asynq.NewServeMux()
	mux.HandleFunc("send_email", utils.HandleSendEmailTask)

	if err := serv.Run(mux); err != nil {
		log.Fatal("Error: Running Mux 'send_email'")
	}

}
