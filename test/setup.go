package test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/your-moon/go-fiber-starter/api"
	"github.com/your-moon/go-fiber-starter/config"
	"github.com/your-moon/go-fiber-starter/models"
	"github.com/your-moon/go-fiber-starter/services"
)

func initAndFreshDB() {
	config.UseTestConfig("your-absolute-path")
	services.InitDB()
	if services.DB == nil {
		panic("DB not initialized")
	}
	services.DB.Migrator().DropTable(&models.User{})
	services.DB.AutoMigrate(&models.User{})
}

func FiberHTTPExpect(t *testing.T) *httpexpect.Expect {
	app := api.Init()
	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(app.Handler()),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})
}
