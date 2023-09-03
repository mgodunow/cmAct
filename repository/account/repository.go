package account

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, account Account) (*Account, error)
	FindByUsername(ctx context.Context, username string) (*Account, error)
	Update(ctx context.Context, account Account) error
}

type mysqlRepository struct {
	mc	*sql.DB
}

func (m *mysqlRepository) Create(ctx context.Context, account Account) (*Account, error) {
	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	_, err = m.mc.Exec("INSERT INTO accounts(username, password, email, bot) VALUES(?, ?, ?, ?)", account.Username, account.PasswordHash, account.Email, false)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (m *mysqlRepository) FindByUsername(ctx context.Context, username string) (*Account, error) {
	
	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	var a Account
	
	row := m.mc.QueryRow("SELECT * FROM accounts WHERE username=?", username)
	err = row.Scan(&a.Username, &a.PasswordHash, &a.Email, &a.Bot)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (m *mysqlRepository) Update(ctx context.Context, account Account) error {
	err := m.mc.Ping()
	if err != nil {
		return err
	}
	_, err = m.mc.Exec("UPDATE accounts SET username = ?, password = ?, email = ?, bot = ? WHERE username = ?",
		account.Username, account.PasswordHash, account.Email, false, account.Username)
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