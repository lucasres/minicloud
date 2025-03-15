package utils

import (
	"context"
	"errors"
)

func GetValueFromCTX[T any](key string, ctx context.Context) (T, error) {
	raw := ctx.Value(key)
	parsed, ok := raw.(T)
	if !ok {
		var zero T
		return zero, errors.New("invalid type for " + key)
	}

	return parsed, nil
}
