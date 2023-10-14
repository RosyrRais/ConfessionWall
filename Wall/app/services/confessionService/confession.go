package confessionservice

import (
	"Wall/app/models"
	"Wall/config/database"
)

func InsertConfession(title string, text string, hidename int, open int, poster string) error {
	confession := models.Confession{
		Title:     title,
		Text:      text,
		Thumbon:   0,
		Thumbdown: 0,
		Hidename:  uint(hidename),
		Open:      uint(open),
		Poster:    poster,
	}
	res := database.DB.Create(&confession)
	return res.Error
}

func CheckConfessionExistById(id int) error {
	res := database.DB.Where("id = ?", id).First(&models.Confession{})
	return res.Error
}

func DeleteConfessionById(id int) error {
	res := database.DB.Where("id = ?", id).Delete(&models.Confession{})
	return res.Error
}

func UpdateConfessionById(id int, title string, text string, hidename int, open int, poster string) error {
	res := database.DB.Save(&models.Confession{Id: uint(id), Title: title, Text: text, Hidename: uint(hidename), Open: uint(open), Poster: poster})
	return res.Error
}

func FindAll() (*[]models.Confession, error) {
	var confessions []models.Confession
	res := database.DB.Find(&confessions)
	if res.Error != nil {
		return nil, res.Error
	}
	return &confessions, nil
}

func FindConfessionsByName(name string) (*[]models.Confession, error) {
	var confessions []models.Confession
	res := database.DB.Where("poster = ?", name).Find(&confessions)
	if res.Error != nil {
		return nil, res.Error
	}
	return &confessions, nil
}
