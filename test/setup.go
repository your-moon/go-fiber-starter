package test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/your-moon/go-fiber-starter/config"
	"github.com/your-moon/go-fiber-starter/internal/api"
	"github.com/your-moon/go-fiber-starter/internal/factory"
	"github.com/your-moon/go-fiber-starter/internal/integrations"
	"github.com/your-moon/go-fiber-starter/internal/models"
	"github.com/your-moon/go-fiber-starter/internal/util/jwtutil"
)

func ActingAs(user models.User, e *httpexpect.Expect) *httpexpect.Expect {

	token := jwtutil.GenerateToken(jwtutil.TokenClaims{
		ID: user.ID,
	})

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	return auth
}

func Auth(e *httpexpect.Expect) *httpexpect.Expect {

	user, err := factory.UserFactory(true, func(u *models.User) {
	})
	if err != nil {
		panic(err)
	}
	token := jwtutil.GenerateToken(jwtutil.TokenClaims{
		ID: user.ID,
	})

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	return auth
}

func initAndFreshDB() {
	config.UseTestConfig("your-absolute-path")
	integrations.InitDB()
	if integrations.DB == nil {
		panic("DB not initialized")
	}
	integrations.DB.Migrator().DropTable(&models.User{})
	integrations.DB.AutoMigrate(&models.User{})
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
