package config

import (
	"database/sql"
	"fmt"

	"cmAct/internal/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func DbConnect() {

	var pass string
	utils.ReadPassDB(&pass)

	fmt.Println(pass)
	connectionString := "cmAct:" + pass + "@tcp(localhost:3306)/cmAct"

	d, err := sql.Open("mysql", connectionString)
	if err != nil {
		logrus.WithError(err).Panic("Error while connecting to db")
	}

	db = d
}

func GetDb() *sql.DB {
	return db
}
