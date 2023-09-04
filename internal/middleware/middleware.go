package middleware

import (
	"cmAct/internal/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func UserIdentity(c *fiber.Ctx) error {
	username, err := jwt.ParseToken(c.Cookies("token"), viper.GetString("SECRET_PHRASE"))
	if err != nil {
		c.Redirect("/")
		return err
	}
	c.Locals("username", username)
	return nil
}
