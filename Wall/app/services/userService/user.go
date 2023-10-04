package userService

import (
	"Wall/app/models"
	"Wall/config/database"
)

func CheckUserExistByName(namee string) error {
	res := database.DB.Where("name = ?", namee).First(&models.User{})
	return res.Error
}

func CheckUserExistByEmail(emaill string) error {
	res := database.DB.Where("email = ?", emaill).First(&models.User{})
	return res.Error
}

func GetUser(namee string) (*models.User, error) {
	var user models.User
	res := database.DB.Where("name = ?", namee).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func ComparePassword(p1 string, p2 string) bool {
	return p1 == p2
}

func InsertUser(name string, email string, sex string, age int, password string) error {
	user := models.User{
		Name:     name,
		Email:    email,
		Sex:      sex,
		Age:      age,
		Password: password,
	}
	res := database.DB.Create(&user)
	return res.Error
}
