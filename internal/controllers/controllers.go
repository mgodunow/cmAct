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

	_, err := models.GetAccountByField(regRequest.Username)
	switch err {
	case sql.ErrNoRows:
		_, err2 := models.GetAccountByField(regRequest.Email)
		if err2 == sql.ErrNoRows {
			var a = models.Account{
				Username: regRequest.Username,
				Email:    regRequest.Email,
				Password: regRequest.Password,
				Bot:      false,
			}
			models.Register(&a)
			return c.Status(fiber.StatusCreated).SendString("Account successfully created")
		}
	case nil:
		return c.Status(fiber.StatusOK).SendString("An account with such data already exists")
	}
	return c.Status(fiber.StatusInternalServerError).SendString("Something went wrong. Please try again")
}
