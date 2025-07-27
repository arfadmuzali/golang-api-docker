package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/hibiken/asynq"
	"gopkg.in/gomail.v2"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func NewEmailPayload(payload EmailPayload) (*asynq.Task, error) {

	data, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return asynq.NewTask("send_email", data), nil
}

func HandleSendEmailTask(ctx context.Context, t *asynq.Task) error {
	var payload EmailPayload

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}
	appPassword := os.Getenv("GOOGLE_APP_PASSWORD")
	log.Println("appPASsword", appPassword)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "arfadmz243@gmail.com", appPassword)
	mail := gomail.NewMessage()
	mail.SetHeader("From", "learnapp@arfad.com")
	mail.SetHeader("To", payload.To)
	mail.SetHeader("Subject", payload.Subject)
	mail.SetBody("text/plain", payload.Body)

	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
