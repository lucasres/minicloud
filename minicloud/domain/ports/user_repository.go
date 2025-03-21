package ports

import (
	"context"

	"lucares.github.com/minicloud/minicloud/domain/entities"
)

const USER_REPOSITORY_KEY_CTX = "userRepoCtx"

type FilterUserOptions struct {
	Limit *int
}

type UserRepositoryPort interface {
	Save(ctx context.Context, u *entities.User) error
	Filter(ctx context.Context, f *FilterUserOptions) ([]*entities.User, error)
}
