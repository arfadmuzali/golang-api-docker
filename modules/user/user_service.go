package user

import (
	"learn/models"
	"learn/utils"
)

func GetAllUser() ([]models.User, error) {
	db := utils.DB

	users := []models.User{}

	result := db.Find(&users)

	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}

func GetUser(id string) (models.User, error) {
	db := utils.DB

	user := models.User{}

	result := db.First(&user, "id = ?", id)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
