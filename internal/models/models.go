package models

import (
	"context"
	"database/sql"

	"cmAct/internal/config"

	"github.com/sirupsen/logrus"
)

var ctx context.Context
var db *sql.DB

// Need a validation on login and registration structures later,
type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegsterRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	Account struct {
		Username string
		Password string
		Email    string
		Bot      bool
	}

	Act struct {
		Username string
		ActName  string
	}

	Task struct {
		Username string
		ActName  string
		Name     string
		Link     string
		//time
	}

	Token struct {
		Token string `json:"token"`
	}
)

func init() {
	config.DbConnect()
	db = config.GetDb()
}

// By default, the user is not connected to the bot, a check for interaction with it will be added later
func Register(a *Account) *Account {
	_, err := db.Exec("INSERT INTO accounts(username, password, email, bot) VALUES(?, ?, ?, ?)", a.Username, a.Password, a.Email, false)
	if err != nil {
		logrus.Warn("Something went wrong while put new account in accounts table", " error: ", err)
	}
	return a
}

func GetAccountByUsername(username string) (*Account, error) {
	var a Account
	row := db.QueryRow("SELECT * FROM accounts WHERE username=?", username)
	err := row.Scan(&a.Username, &a.Password, &a.Email, &a.Bot)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Account{}, err
		}
		logrus.Warn("Error while scanning the rows", " error: ", err)
		return &Account{}, err
	}
	return &a, nil
}

func GetAccountByEmail(email string) (*Account, error) {
	var a Account
	row := db.QueryRow("SELECT * FROM accounts WHERE email=?", email)
	err := row.Scan(&a.Username, &a.Password, &a.Email, &a.Bot)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Account{}, err
		}
		logrus.Warn("Error while scanning the rows", " error: ", err)
		return &Account{}, err
	}
	return &a, nil
}
