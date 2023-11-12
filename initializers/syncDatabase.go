package initializers

import "github.com/vnaxel/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
}