package factory

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/your-moon/go-fiber-starter/models"
	"github.com/your-moon/go-fiber-starter/services"
)

func userFactory(commitDB bool) models.User {
	var user models.User
	gofakeit.Struct(&user)

	if commitDB {
		services.DB.Create(&user)
	}

	return user
}
