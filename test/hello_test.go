package test

import (
	"testing"

	"github.com/your-moon/go-fiber-starter/internal/factory"
)

func TestHello(t *testing.T) {
	initAndFreshDB()

	e := FiberHTTPExpect(t)
	e.GET("/").Expect().Status(200).Text().IsEqual("Hello, World!")
}

func TestGetUser(t *testing.T) {
	initAndFreshDB()

	user := factory.UserFactory(true)

	e := FiberHTTPExpect(t)
	e.GET("/user/1").Expect().Status(200).JSON().Object().ContainsSubset(map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	})
}

func TestCreateUser(t *testing.T) {
	initAndFreshDB()

	e := FiberHTTPExpect(t)
	e.POST("/user").WithJSON(map[string]interface{}{
		"email": "tester@gmail.com",
	}).Expect().Status(200).JSON().Object().ContainsSubset(map[string]interface{}{
		"email": "tester@gmail.com",
	})

}
