package controllers

import (
	"database/sql"
	"context"

	"cmAct/internal/hash"
	"cmAct/internal/jwt"
	"cmAct/internal/utils"
	"cmAct/internal/repository/user"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TODO: check on bot activity and encryption of password.
// Also need to catch answers from server on client side and make an popups with pushes(Success register, or conflict, or 502).
// Register need a refactor later.
type RegisterRequest struct{
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type LoginRequest struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

type accRepository struct {
	repository	user.Repository
}

func (a *accRepository) Register(c *fiber.Ctx) error {
	var regRequest RegisterRequest
	if err := c.BodyParser(&regRequest); err != nil {
		logrus.Info(fiber.ErrBadRequest, " error: ", err)
		return err
	}
	
	u, err := a.repository.FindByUsername(context.Background(),regRequest.Username)
	switch err {
		case sql.ErrNoRows:
			valid := utils.RegisterValidate(u.Username, u.Email, u.PasswordHash)
			if !valid {
				return c.Status(fiber.StatusOK).SendString("Invalid data for registration. Please try again")
			}

			a.repository.Create(context.Background(), *u)

			return c.Status(fiber.StatusCreated).SendString("Account successfully created")

		default:
			return c.Status(fiber.StatusOK).SendString("An account with such data already exists")
	}
}

//TODO: refactor login
func (a *accRepository) Login(c *fiber.Ctx) error {
	var loginRequest LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		logrus.Info(fiber.ErrBadRequest, " error: ", err)
		return err
	}
	u, err := a.repository.FindByUsername(context.Background(),loginRequest.Username)
	if err != nil || u == nil { //May be u == nil is mistake matching
		return c.Status(fiber.StatusUnauthorized).SendString("There is no account with provided email")
	}
	if u.PasswordHash != hash.Hash(viper.GetString("SALT"), loginRequest.Password) {
		return c.Status(fiber.StatusUnauthorized).SendString("Wrong password. Please, try again")
	}

	token, err := jwt.GenerateToken(u.Email, viper.GetString("SECRET_PHRASE"))
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(token)
	}

	return c.Status(fiber.StatusUnauthorized).SendString("Authorization error. Please, try again")
}

//TODO: refactor home

// func Home(c *fiber.Ctx) error {
// 	username, err := jwt.ParseToken(c.Cookies("token"), viper.GetString("SECRET_PHRASE"))
// 	if err != nil {
// 		return err
// 	}
// 	return c.Render("home", fiber.Map{
// 		"username": username,
// 	})
// }

func NewRegisterLoginController(accRepo user.Repository) *accRepository{
	return &accRepository{repository: accRepo}
}
