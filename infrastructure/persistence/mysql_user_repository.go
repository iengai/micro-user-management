package persistence

import (
	"database/sql"
	"fmt"
	"micro-user-management/internal/domain/entity"
	"micro-user-management/internal/domain/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", user.ID, user.Username, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (r *UserRepository) Find(id int64) (*entity.User, error) {
	var user entity.User
	row := r.db.QueryRow("SELECT id, username, password FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Username, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, &errors.NotFoundError{}
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	row := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	if err := row.Scan(&user.ID, &user.Username, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, &errors.NotFoundError{}
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	return &user, nil
}
