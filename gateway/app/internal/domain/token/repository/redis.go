package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/go-redis/redis"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("token-repo")
)

type repo struct {
	client redis.UniversalClient
}

func NewTokenRepo(client redis.UniversalClient) *repo {
	return &repo{client: client}
}

func (t *repo) Set(ctx context.Context, refreshToken string, userId int, expireAt time.Duration) error {
	ctx, span := tracer.Start(ctx, "repo.Set")
	defer span.End()
	logger := logging.LoggerFromContext(ctx)
	logger.Sugar().Debugf("try to save token = %v for user with id = %v", refreshToken, userId)
	err := t.client.Set(refreshToken, strconv.Itoa(userId), expireAt).Err()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (t *repo) Get(ctx context.Context, refreshToken string) (int, error) {
	ctx, span := tracer.Start(ctx, "repo.Get")
	defer span.End()
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

func (t *repo) Delete(ctx context.Context, refreshToken string) error {
	ctx, span := tracer.Start(ctx, "repo.Delete")
	defer span.End()
	err := t.client.Del(refreshToken).Err()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
