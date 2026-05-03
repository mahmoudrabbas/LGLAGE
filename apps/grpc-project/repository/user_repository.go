package repository

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUserById(id int64) (int64, string, string, error) {
	row := r.DB.QueryRow("SELECT id, name, email FROM users where id=?", id)

	var userId int64
	var name, email string

	err := row.Scan(&userId, &name, &email)

	return userId, name, email, err
}

func (r *UserRepository) CreateUser(name, email string) (int64, string, string, error) {
	res, err := r.DB.Exec("INSERT INTO users (name, email) VALUES(?,?)", name, email)

	if err != nil {
		return 0, "", "", err
	}

	id, err := res.LastInsertId()

	return id, name, email, err
}

func (r *UserRepository) UpdateUser(id int64, name, email string) (int64, string, string, error) {

	res, err := r.DB.Exec("UPDATE users SET name=?, email=? WHERE id=?", name, email, id)

	if err != nil {
		return 0, "", "", err
	}

	rows, _ := res.RowsAffected()

	if rows == 0 {
		return 0, "", "", sql.ErrNoRows
	}

	return id, name, email, nil

}

func (r *UserRepository) DeleteUser(id int64) (int64, error) {
	res, err := r.DB.Exec("DELETE FROM users where id=?", id)

	if err != nil {
		return 0, err
	}

	row, err := res.RowsAffected()

	return row, err
}
