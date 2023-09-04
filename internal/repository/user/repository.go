package user

import (
	"context"
	"database/sql"
)
//TODO: FindByEmail

type Repository interface {
	Create(ctx context.Context, user User) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Update(ctx context.Context, user User) error
}

type mysqlRepository struct {
	mc	*sql.DB
}

func (m *mysqlRepository) Create(ctx context.Context, user User) (*User, error) {
	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	_, err = m.mc.Exec("INSERT INTO users(username, password, email, bot) VALUES(?, ?, ?, ?)", user.Username, user.PasswordHash, user.Email, false)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *mysqlRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
	
	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	var a User
	
	row := m.mc.QueryRow("SELECT * FROM users WHERE username=?", username)
	err = row.Scan(&a.Username, &a.PasswordHash, &a.Email, &a.Bot)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (m *mysqlRepository) Update(ctx context.Context, user User) error {
	err := m.mc.Ping()
	if err != nil {
		return err
	}
	_, err = m.mc.Exec("UPDATE users SET username = ?, password = ?, email = ?, bot = ? WHERE username = ?",
		user.Username, user.PasswordHash, user.Email, false, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &mysqlRepository{
		mc: db,
	}
}