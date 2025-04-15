package userusecase

import (
	"context"

	userdomain "github.com/lunatictiol/clean-architecture-in-go/internal/services/user/userDomain"
	userrepository "github.com/lunatictiol/clean-architecture-in-go/internal/services/user/userRepository"
)

type UserUsecase struct {
	repo *userrepository.UserRepository
}

func NewUserUsecase(r *userrepository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id int) (*userdomain.User, error) {
	return uc.repo.FindByID(ctx, id)
}
