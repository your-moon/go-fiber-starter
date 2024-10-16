package test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/your-moon/go-fiber-starter/api"
)

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

func TestHello(t *testing.T) {
	freshDB()

	e := FiberHTTPExpect(t)
	if e == nil {
		t.Fatal("e is nil")
	}

	e.GET("/").Expect().Status(200).Text().IsEqual("Hello, World!")
}
