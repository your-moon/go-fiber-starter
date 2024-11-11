package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/your-moon/go-fiber-starter/internal/models"
	"github.com/your-moon/go-fiber-starter/internal/services"
)

func UserFactory(commitDB bool) models.User {
	var user models.User
	gofakeit.Struct(&user)

	if commitDB {
		services.DB.Create(&user)
	}

	return user
}
