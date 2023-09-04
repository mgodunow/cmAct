package config

import (
	"context"
	"database/sql"
	"log"

	"cmAct/internal/controllers"
	"cmAct/internal/repository/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)
type App interface {
	ListenAndServe() error
}

type app struct {
	mc		*sql.DB
	ctx		context.Context
	config *Config
}

func (a *app) ListenAndServe() error {
	//Engine & server initializing
	
	engine := html.New("../templates", ".html")

	server := fiber.New(fiber.Config{
		Views: engine,
	})
	
	log.Println("App is running")

	//
	c := controllers.NewRegisterLoginController(user.NewRepository(a.mc))
	//Register controllers. May be register them in other function later

	// server.Get("/home", controllers.Home)
	server.Static("/", "../static")
	server.Post("/login", c.Login)
	server.Post("/registration", c.Register)

	//localhost serve
	log.Fatal(server.Listen(":8080"))

	return nil
}

func New() (*app, error) {
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	mc, err := sql.Open("mysql", "cmAct:" + config.DbPass + "@tcp(localhost:3306)/cmAct")
	if err != nil {
		log.Panic(err)
	}

	return &app{
		config: &config,
		mc: mc,
	}, nil

}