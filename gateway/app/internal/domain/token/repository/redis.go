package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/go-redis/redis"
)

type tokenStorage struct {
	client redis.UniversalClient
}

func NewChoiceCache(client redis.UniversalClient) *tokenStorage {
	return &tokenStorage{client: client}
}

func (t *tokenStorage) Set(ctx context.Context, refreshToken string, userId int, expireAt time.Duration) error {
	logger := logging.LoggerFromContext(ctx)
	logger.Sugar().Debugf("try to save token = %v for user with id = %v", refreshToken, userId)
	err := t.client.Set(refreshToken, strconv.Itoa(userId), expireAt).Err()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (t *tokenStorage) Get(ctx context.Context, refreshToken string) (int, error) {
	value, err := t.client.Get(refreshToken).Result()
	if err != nil {
		return -1, errors.New(err.Error())
	}
	count, err := strconv.Atoi(value)
	if err != nil {
		return -1, errors.Wrap(err, "couldn't parse user id")
	}
	return count, nil
}

func (t *tokenStorage) Delete(ctx context.Context, refreshToken string) error {
	err := t.client.Del(refreshToken).Err()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
