package controllers

import (
	"database/sql"

	"cmAct/internal/hash"
	jwtGenerate "cmAct/internal/jwt"
	"cmAct/internal/models"
	"cmAct/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Later will add check on bot activity and encryption of password.
// Also need to catch answers from server on client side and make an popups with pushes(Success register, or conflict, or 502).
// Register need a refactor later.
func Register(c *fiber.Ctx) error {
	regRequest := models.RegsterRequest{}
	if err := c.BodyParser(&regRequest); err != nil {
		logrus.Info(fiber.ErrBadRequest, " error: ", err)
		return err
	}

	_, err := models.GetAccountByUsername(regRequest.Username)
	switch err {
	case sql.ErrNoRows:
		_, err2 := models.GetAccountByEmail(regRequest.Email)
		switch err2 {
		case sql.ErrNoRows:
			var a = models.Account{
				Username: regRequest.Username,
				Email:    regRequest.Email,
				Password: regRequest.Password,
				Bot:      false,
			}
			valid := utils.RegisterValidate(a.Username, a.Email, a.Password)
			if !valid {
				return c.Status(fiber.StatusOK).SendString("Invalid data for registration is indicated. Please try again")
			}
			models.Register(&a)
			return c.Status(fiber.StatusCreated).SendString("Account successfully created")
		default:
			return c.Status(fiber.StatusOK).SendString("An account with such data already exists")
		}
	default:
		return c.Status(fiber.StatusOK).SendString("An account with such data already exists")
	}
}

func Login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		logrus.Info(fiber.ErrBadRequest, " error: ", err)
		return err
	}
	a, err := models.GetAccountByEmail(loginRequest.Email)
	if err != nil || (*a == models.Account{}) {
		return c.Status(fiber.StatusUnauthorized).SendString("There is no account with provided email")
	}
	if a.Password != hash.Hash(loginRequest.Password) {
		return c.Status(fiber.StatusUnauthorized).SendString("Wrong password. Please, try again")
	}

	token, err := jwtGenerate.GenerateToken(a.Email)
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(models.Token{Token: token})
	}

	return c.Status(fiber.StatusUnauthorized).SendString("Authorization error. Please, try again")
}
