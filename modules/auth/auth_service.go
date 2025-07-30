package auth

import (
	"learn/models"
	"learn/utils"
	"os"

	"github.com/hibiken/asynq"
	"golang.org/x/crypto/bcrypt"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func HandleLogin(body *LoginDto) (TokenResponse, error) {

	db := utils.DB

	user := models.User{Email: body.Email}

	result := db.First(&user, "email = ?", body.Email)

	if result.Error != nil {
		return TokenResponse{}, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return TokenResponse{}, err
	}

	redisOpt := asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")}

	queue := asynq.NewClient(redisOpt)
	defer queue.Close()

	task, err := utils.NewEmailPayload(utils.EmailPayload{To: user.Email, Subject: "Email Confirmation", Body: "this is email confimation, you logedin on x platform"})

	queue.Enqueue(task)

	token, err := utils.GenerateToken(user.Name)

	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{Token: token}, err
}

func HandleRegister(body *RegisterDto) (models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	db := utils.DB

	user := models.User{Name: body.Username, Password: string(hashedPassword), Email: body.Email}

	result := db.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil

}
