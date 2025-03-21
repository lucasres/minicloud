package users

import (
	"context"

	"lucares.github.com/minicloud/minicloud/domain/entities"
	"lucares.github.com/minicloud/minicloud/domain/ports"
	"lucares.github.com/minicloud/minicloud/shared/utils"
)

var instanceUC *CreateUserUseCase

type CreateUserUseCase struct {
	repo ports.UserRepositoryPort
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, u *entities.User) error {
	return uc.Execute(ctx, u)
}

func NewWasConfiguredUseCase(ctx context.Context) (*CreateUserUseCase, error) {
	if instanceUC == nil {
		repo, err := utils.GetValueFromCTX[ports.UserRepositoryPort](ports.USER_REPOSITORY_KEY_CTX, ctx)
		if err != nil {
			return nil, err
		}

		instanceUC = &CreateUserUseCase{
			repo: repo,
		}
	}

	return instanceUC, nil
}
