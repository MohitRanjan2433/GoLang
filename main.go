package main

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

type User struct{
	UserName   string     `json:"username"`
	Role       string     `json:"role"`
}

func main() {
	app := fiber.New()

	app.Get("/post", handleGetPost)
	app.Get("/post/manage", onlyAdmin(handleGetPostManage))

	log.Fatal(app.Listen(":4000"))
}

func onlyAdmin(fn fiber.Handler) fiber.Handler{
	return func (c *fiber.Ctx) error {
		user := getUserFromDB()
		if user.Role != "admin" {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return fn(c)
	}
}

func getUserFromDB() User {
	return User{
		UserName: "Mohit2433",
		Role: "admin",
	}
}

func handleGetPost(c *fiber.Ctx) error {
	return c.JSON("some posts here")
}

func handleGetPostManage(c *fiber.Ctx) error {
	return c.JSON("Some admin posts are here")
}