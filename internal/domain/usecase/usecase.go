package usecase

import (
	"errors"
	"micro-user-management/internal/domain/entity"
	domainErrors "micro-user-management/internal/domain/errors"
	"micro-user-management/internal/domain/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uc *UserUsecase) Register(user *entity.User) error {
	if !user.IsValid() {
		return errors.New("Invalid user data")
	}

	existingUser, err := uc.userRepo.FindByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, &domainErrors.NotFoundError{}) {
			return err
		}
	}
	if existingUser != nil {
		return errors.New("User already exists")
	}

	if err := uc.userRepo.Save(user); err != nil {
		return err
	}

	return nil
}

func (uc *UserUsecase) Login(user *entity.User) (*entity.User, error) {
	if !user.IsValid() {
		return nil, errors.New("Invalid user data")
	}
	existingUser, err := uc.userRepo.Find(user.ID)
	if err != nil {
		return nil, err
	}

	if existingUser.PasswordHash != user.PasswordHash {
		return nil, errors.New("Invalid password")
	}

	return existingUser, nil
}

func (uc *UserUsecase) GetUser(id int64) (*entity.User, error) {
	user, err := uc.userRepo.Find(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
