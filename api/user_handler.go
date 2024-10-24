package api

import (
	"github.com/AshrithSathu/htelapi/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Foo",
	}
	return c.JSON(u)

}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "James Foo",
	})
}
