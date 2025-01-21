package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/your-moon/go-fiber-starter/internal/integrations"
	"github.com/your-moon/go-fiber-starter/internal/models"
)

type OptionUser func(*models.User)

func UserFactory(commitDB bool, opts ...OptionUser) (models.User, error) {
	var user models.User
	err := gofakeit.Struct(&user)
	if err != nil {
		return models.User{}, err
	}

	for _, opt := range opts {
		opt(&user)
	}

	if commitDB {
		integrations.DB.Create(&user)
	}

	return user, nil
}
