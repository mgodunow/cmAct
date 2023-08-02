package middleware

import (
	"cmAct/internal/jwt"

	"github.com/gofiber/fiber/v2"
)

func UserIdentity(c *fiber.Ctx) error {
	username, err := jwt.ParseToken(c.Cookies("token"))
	if err != nil {
		c.Redirect("/")
		return err
	}
	c.Locals("username", username)
	return nil
}
