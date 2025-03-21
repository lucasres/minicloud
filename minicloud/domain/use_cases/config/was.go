package config

import (
	"context"

	"lucares.github.com/minicloud/minicloud/domain/ports"
	"lucares.github.com/minicloud/minicloud/shared/utils"
)

var instance *WasConfiguredUseCase

type WasConfiguredUseCase struct {
	repo ports.UserRepositoryPort
}

func (wc *WasConfiguredUseCase) Execute(ctx context.Context) (bool, error) {
	limit := 1
	f := &ports.FilterUserOptions{
		Limit: &limit,
	}

	rs, err := wc.repo.Filter(ctx, f)
	if err != nil {
		return false, err
	}

	return len(rs) == 1, nil
}

func NewWasConfiguredUseCase(ctx context.Context) (*WasConfiguredUseCase, error) {
	if instance == nil {
		repo, err := utils.GetValueFromCTX[ports.UserRepositoryPort](ports.USER_REPOSITORY_KEY_CTX, ctx)
		if err != nil {
			return nil, err
		}

		instance = &WasConfiguredUseCase{
			repo: repo,
		}
	}

	return instance, nil
}
