package user

import (
	"context"
	"database/sql"

	"github.com/TimurZheksimbaev/Golang-webchat/utils"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository  {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	var lastInsertedID int
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertedID)
	if err != nil {
		return &User{}, utils.DatabaseError("Could not insert user", err)
	}
	user.ID = int64(lastInsertedID)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query :=  "SELECT id, username, password, email FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		return &User{}, utils.DatabaseError("Could not get user by email", err)
	}

	return &u, nil
}