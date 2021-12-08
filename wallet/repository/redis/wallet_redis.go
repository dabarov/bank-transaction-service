package redis

import (
	"time"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/go-redis/redis"
)

type walletRedisRepository struct {
	redisClient *redis.Client
	timeout     time.Duration
	secret      string
}

func NewWalletRedisRepository(redisClient *redis.Client, timeout int, secret string) domain.WalletRedisRepository {
	return &walletRedisRepository{
		redisClient: redisClient,
		timeout:     time.Duration(timeout) * time.Second,
		secret:      secret,
	}
}

func (w *walletRedisRepository) GetValue(key string) (string, error) {
	return w.redisClient.Get(key).Result()
}

func (w *walletRedisRepository) GetSecret() string {
	return w.secret
}
