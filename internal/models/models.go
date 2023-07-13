package models

import (
	"database/sql"

	"cmAct/internal/config"
)

var db *sql.DB

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
		Email    string
		Password string
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
)

func init() {
	config.DbConnect()
	db = config.GetDb()
}
