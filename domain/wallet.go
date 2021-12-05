package domain

import "context"

type Wallet struct {
	ID                  uint64 `json:"id"`
	User                User   `json:"user"`
	LastTransactionDate string `json:"last_transaction_date"`
	Balance             string `json:"balance"`
	CreatedAt           string `json:"created_at"`
}

type WalletUsecase interface {
	Create(ctx context.Context, iin string) error
	Deposit(ctx context.Context, walletID uint64, amount uint64) error
	Transfer(ctx context.Context, walletFromID uint64, walletToID uint64, amount uint64) error
	GetUserWallets(ctx context.Context, iin string) ([]*Wallet, error)
}

type WalletRepository interface {
	Create(ctx context.Context, iin string) error
	Deposit(ctx context.Context, walletID uint64, amount uint64) error
	Transfer(ctx context.Context, walletFromID uint64, walletToID uint64, amount uint64) error
	GetUserWallets(ctx context.Context, iin string) ([]*Wallet, error)
}
