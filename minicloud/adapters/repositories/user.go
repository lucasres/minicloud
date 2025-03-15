package repositories

import (
	"context"
	"database/sql"

	"lucares.github.com/minicloud/minicloud/domain/entities"
	"lucares.github.com/minicloud/minicloud/domain/ports"
)

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) Update(ctx context.Context, u *entities.User) {

}

func (ur *UserRepository) Filter(ctx context.Context, f *ports.FilterUserOptions) ([]*entities.User, error) {
	return nil, nil
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
