package auth

import (
	"fmt"
	"learn/utils"
	"os"

	"github.com/hibiken/asynq"
	"golang.org/x/crypto/bcrypt"
)

type TokenResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HandleLogin(body *LoginDto) (TokenResponse, error) {
	if body.Username != "arfad" || body.Password != "beraktakcebok" {
		return TokenResponse{}, fmt.Errorf("Username or Password are wrong")
	}

	redisOpt := asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")}

	queue := asynq.NewClient(redisOpt)
	defer queue.Close()

	task, err := utils.NewEmailPayload(utils.EmailPayload{To: "arfadmuzali258@gmail.com", Subject: "Email Confirmation", Body: "this is email confimation, you logedin on x platform"})

	queue.Enqueue(task)

	token, err := utils.GenerateToken(body.Username)

	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{Token: token}, err
}

func HandleRegister(body *RegisterDto) (RegisterResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{UserId: "abc123", Email: body.Email, Password: string(hashedPassword), Username: body.Username}, nil

}
