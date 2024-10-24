package main

import (
	"flag"

	"github.com/AshrithSathu/htelapi/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4000", "server listen address")
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "Foo",
	})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "James Foo",
	})
}
