package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lucasres/goquent"
	"lucares.github.com/minicloud/minicloud/domain/entities"
	"lucares.github.com/minicloud/minicloud/domain/ports"
)

type UserRepository struct {
	db        *sql.DB
	tableName string
	fields    []string
}

func (ur *UserRepository) Save(ctx context.Context, u *entities.User) error {
	return nil
}

func (ur *UserRepository) Filter(ctx context.Context, f *ports.FilterUserOptions) ([]*entities.User, error) {
	q := goquent.New(goquent.PGSQL).
		Select().
		From(ur.tableName)

	if f.Limit != nil {
		q.Limit(*f.Limit)
	}

	result := make([]*entities.User, 0)
	err := goquent.QueryContext(ctx, ur.db, q, func(rows *sql.Rows) error {
		var u *entities.User
		err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.Password,
			&u.Name,
		)
		if err != nil {
			return fmt.Errorf("cant scan user row: %w", err)
		}
		result = append(result, u)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("cant query user filter: %w", err)
	}

	return result, nil
}

func NewRepository(db *sql.DB) ports.UserRepositoryPort {
	return &UserRepository{
		db:        db,
		tableName: "users",
		fields:    []string{"id", "email", "password", "name"},
	}
}
