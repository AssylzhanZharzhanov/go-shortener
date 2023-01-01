package repository

import (
	"context"
	"github.com/google/uuid"
	"log"
	"testing"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
)

func SetupRedis() domain.ShortenerRedisRepository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	authRedisRepo := NewRedisRepository(client)
	return authRedisRepo
}

func TestRedisRepository_Set(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
	)

	redisRepoStub := SetupRedis()

	t.Run("Set", func(t *testing.T) {
		value := &domain.Link{
			ID:          1,
			OriginalURL: "test",
			ShortURL:    "test",
		}

		err := redisRepoStub.Set(ctx, value.ShortURL, value, 10)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

func TestRedisRepository_Get(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
	)

	redisRepoStub := SetupRedis()

	t.Run("Get", func(t *testing.T) {
		u := &domain.Link{
			ID:          1,
			OriginalURL: "test",
			ShortURL:    "test",
		}

		err := redisRepoStub.Set(ctx, u.ShortURL, u, 10)
		require.NoError(t, err)
		require.Nil(t, err)

		user, err := redisRepoStub.Get(ctx, u.ShortURL)
		require.NoError(t, err)
		require.NotNil(t, user)
	})
}

func TestRedisRepository_Delete(t *testing.T) {
	t.Parallel()

	redisRepoStub := SetupRedis()

	t.Run("Delete", func(t *testing.T) {
		key := uuid.New().String()

		err := redisRepoStub.Delete(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}
