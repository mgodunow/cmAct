package act 

import (
	"database/sql"
	"context"
)

type Repository interface {
	Create(ctx context.Context, act Act) (*Act, error)
	FindOne(ctx context.Context, actname string) (*Act, error)
	FindAll(ctx context.Context, username string) (*[]Act, error)
	Delete(ctx context.Context, actname string) error
}

type mysqlRepository struct {
	mc	*sql.DB
}

func (m *mysqlRepository) Create(ctx context.Context, act Act) (*Act, error) {

	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	_, err = m.mc.Exec("INSERT INTO acts(username,	actname) VALUES(?, ?)", act.Username, act.ActName)
	if err != nil {
		return nil, err
	}
	return &act, nil
}

func (m *mysqlRepository) FindOne(ctx context.Context, actname string) (*Act, error) {

	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	var a Act

	row := m.mc.QueryRow("SELECT * FROM acts WHERE actName = ?", actname)
	err = row.Scan(&a.Username, &a.ActName)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (m *mysqlRepository) FindAll(ctx context.Context, username string) (*[]Act, error) {
	var a *Act
	
	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	acts := make([]Act, 0,)
	rows, err := m.mc.Query("SELECT * FROM acts WHERE actName = ?", username)
	if err != nil {
		return nil, err
	}

	for {

		if !rows.Next() {
			break
		}

		err = rows.Scan(&a.Username, &a.ActName)

		if err != nil {
			return &acts, err
		}

		acts = append(acts, *a)

	}
	
	return &acts, nil
}

func (m *mysqlRepository) Delete(ctx context.Context, actName string) error {

	err := m.mc.Ping()
	if err != nil {
		return err
	}

	_, err = m.mc.Exec("DELETE FROM acts WHERE actName = ?", actName)
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