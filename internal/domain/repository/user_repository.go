package repository

import (
	"micro-user-management/internal/domain/entity"
)

type UserRepository interface {
	Save(user *entity.User) error
	Find(id int64) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
}
