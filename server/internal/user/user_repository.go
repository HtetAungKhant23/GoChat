package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (repo *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var newUser User
	query := `INSERT INTO users(id, username, email, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`

	err := repo.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.UserName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&newUser.ID,
		&newUser.UserName,
		&newUser.Email,
		&newUser.Password,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)

	if err != nil {
		return &User{}, err
	}

	return &newUser, nil
}
