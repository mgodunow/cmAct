package controllers

import (
	"database/sql"

	"cmAct/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Later will add check on bot activity and encryption of password.
// Also need to catch answers from server on client side and make an popups with pushes(Success register, or conflict, or 502).
func Register(c *fiber.Ctx) error {
	regRequest := models.RegsterRequest{}
	if err := c.BodyParser(&regRequest); err != nil {
		logrus.Info(fiber.ErrBadRequest, " error: ", err)
		return err
	}

	_, err := models.GetAccount(regRequest.Username)
	switch err {
	case sql.ErrNoRows:
		var a = models.Account{
			Username: regRequest.Username,
			Email:    regRequest.Email,
			Password: regRequest.Password,
			Bot:      false,
		}
		models.Register(&a)
		c.SendStatus(fiber.StatusCreated)
		return c.Redirect("/")
	case nil:
		logrus.Info("An account with the same email or username already exists")
		c.SendStatus(fiber.StatusConflict)
		return c.Redirect("/")
	}
	return c.SendStatus(fiber.StatusInternalServerError)
}
