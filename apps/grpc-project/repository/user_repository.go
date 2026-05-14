package repository

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name, email string) (int64, string, string, error)
	GetUserById(ctx context.Context, id int64) (int64, string, string, error)
	UpdateUser(ctx context.Context, id int64, name, email string) (int64, string, string, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
}

type MysqlUserRepository struct {
	DB *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) UserRepository {
	return &MysqlUserRepository{
		DB: db,
	}
}

func (r *MysqlUserRepository) GetUserById(ctx context.Context, id int64) (int64, string, string, error) {
	row := r.DB.QueryRowContext(ctx, "SELECT id, name, email FROM users where id=?", id)

	var userId int64
	var name, email string

	err := row.Scan(&userId, &name, &email)

	return userId, name, email, err
}

func (r *MysqlUserRepository) CreateUser(ctx context.Context, name, email string) (int64, string, string, error) {
	res, err := r.DB.ExecContext(ctx, "INSERT INTO users (name, email) VALUES(?,?)", name, email)

	if err != nil {
		return 0, "", "", err
	}

	id, err := res.LastInsertId()

	return id, name, email, err
}

func (r *MysqlUserRepository) UpdateUser(ctx context.Context, id int64, name, email string) (int64, string, string, error) {

	res, err := r.DB.ExecContext(ctx, "UPDATE users SET name=?, email=? WHERE id=?", name, email, id)

	if err != nil {
		return 0, "", "", err
	}

	rows, _ := res.RowsAffected()

	if rows == 0 {
		return 0, "", "", sql.ErrNoRows
	}

	return id, name, email, nil

}

func (r *MysqlUserRepository) DeleteUser(ctx context.Context, id int64) (int64, error) {
	res, err := r.DB.ExecContext(ctx, "DELETE FROM users where id=?", id)

	if err != nil {
		return 0, err
	}

	row, err := res.RowsAffected()

	return row, err
}
