package redis

import (
	"context"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/go-redis/redis/v8"
)

type walletRedisRepository struct {
	redisClient *redis.Client
}

func NewWalletRedisRepository(redisClient *redis.Client) domain.WalletRepository {
	return &walletRedisRepository{redisClient}
}

func (w *walletRedisRepository) Create(ctx context.Context, iin string) error {
	return nil
}

func (w *walletRedisRepository) Deposit(ctx context.Context, wallet *domain.Wallet, amount uint64) error {
	return nil
}

func (w *walletRedisRepository) Transfer(ctx context.Context, walletFrom *domain.Wallet, walletTo *domain.Wallet, amount uint64) error {
	return nil
}

func (w *walletRedisRepository) GetUserWallets(ctx context.Context, user *domain.User) ([]*domain.Wallet, error) {
	return nil, nil
}
