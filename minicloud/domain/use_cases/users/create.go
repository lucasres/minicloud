package users

import (
	"context"

	"lucares.github.com/minicloud/minicloud/domain/entities"
)

type CreateUserUseCase struct{}

func (*CreateUserUseCase) Execute(ctx context.Context, u *entities.User) error {

	return nil
}
