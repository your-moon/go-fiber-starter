package test

import (
	"github.com/your-moon/go-fiber-starter/models"
	"github.com/your-moon/go-fiber-starter/services"
)

func freshDB() {
	services.InitDB()
	if services.DB == nil {
		panic("DB not initialized")
	}
	services.DB.Migrator().DropTable(&models.User{})
	services.DB.AutoMigrate(&models.User{})
}
