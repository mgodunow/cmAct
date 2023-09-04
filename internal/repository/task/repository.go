package task

import (
	"database/sql"
	"context"
)

type Repository interface {
	Create(ctx context.Context, task Task) (*Task, error)
	Find(ctx context.Context, username, actname, taskname string) (*Task, error)
	FindAll(ctx context.Context, username, actname string) ([]Task, error)
	Delete(ctx context.Context, taskname string) error
	DeleteAll(ctx context.Context, username, actname string) error
}

type mysqlRepository struct {
	mc	*sql.DB
}

func (m *mysqlRepository) Create(ctx context.Context, task Task) (*Task, error) {

	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	_, err = m.mc.Exec("INSERT INTO tasks(username,	actname, taskname, link, datetime) VALUES(?, ?, ?, ?, ?)",
		task.Username, task.ActName, task.TaskName, task.Link, task.DateTime)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (m *mysqlRepository) Find(ctx context.Context, username, actname, taskname string) (*Task, error) {
	
	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	var t Task

	row := m.mc.QueryRow("SELECT * FROM tasks WHERE username = ? AND actname = ? AND taskname = ?",
		username, actname, taskname)
	err = row.Scan(&t.Username, &t.ActName, &t.TaskName, &t.Link, &t.DateTime)
		if err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *mysqlRepository) FindAll(ctx context.Context, username, actname string) ([]Task, error) {
	
	err := m.mc.Ping()
	if err != nil {
		return nil, err
	}

	var t *Task
	tasks := make([]Task, 0)

	rows, err := m.mc.Query("SELECT * FROM tasks WHERE username = ? AND actname = ?", username, actname)
	if err != nil {
		return nil, err
	}

	for {

		if !rows.Next() {
			break
		}

		err = rows.Scan(&t.Username, &t.ActName, &t.TaskName, &t.Link, &t.DateTime)

		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, *t)

	}
	
	return tasks, nil
}

func (m *mysqlRepository) Delete(ctx context.Context, taskname string) error {
	
	err := m.mc.Ping()
	if err != nil {
		return err
	}

	_, err = m.mc.Exec("DELETE FROM tasks WHERE taskname = ?", taskname)
	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlRepository) DeleteAll(ctx context.Context, username, actname string) error {
	
	err := m.mc.Ping()
	if err != nil {
		return err
	}

	_, err = m.mc.Exec("DELETE FROM tasks WHERE username = ? AND actname = ?", username, actname)

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