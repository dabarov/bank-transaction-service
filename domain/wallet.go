package domain

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	ID                  uuid.UUID `json:"id"`
	UserIIN             string    `json:"user"`
	LastTransactionDate string    `json:"last_transaction_date"`
	Balance             uint64    `json:"balance"`
	CreatedAt           string    `json:"created_at"`
}

func (wallet *Wallet) BeforeCreate(tx *gorm.DB) (err error) {
	wallet.ID, err = uuid.NewUUID()
	return err
}

type WalletUsecase interface {
	Create(ctx context.Context, iin string) error
	Deposit(ctx context.Context, iin string, walletID uuid.UUID, amount uint64) error
	Transfer(ctx context.Context, iin string, walletFromID uuid.UUID, walletToID uuid.UUID, amount uint64) error
	GetUserWallets(ctx context.Context, iin string) ([]byte, error)
	GetRedisValue(key string) (string, error)
	GetRedisSecret() string
}

type WalletDBRepository interface {
	Create(ctx context.Context, iin string) error
	Deposit(ctx context.Context, iin string, walletID uuid.UUID, amount uint64) error
	Transfer(ctx context.Context, iin string, walletFromID uuid.UUID, walletToID uuid.UUID, amount uint64) error
	GetUserWallets(ctx context.Context, iin string) ([]Wallet, error)
}

type WalletRedisRepository interface {
	GetValue(key string) (string, error)
	GetSecret() string
}
