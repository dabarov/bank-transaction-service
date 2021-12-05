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
	Create(ctx context.Context, iin uint64) error
	Deposit(ctx context.Context, wallet *Wallet, amount uint64) error
	Transfer(ctx context.Context, walletFrom *Wallet, walletTo *Wallet, amount uint64) error
	GetUserWallets(ctx context.Context, user *User) ([]*Wallet, error)
}

type WalletRepository interface {
	Create(ctx context.Context, iin uint64) error
	Deposit(ctx context.Context, wallet *Wallet, amount uint64) error
	Transfer(ctx context.Context, walletFrom *Wallet, walletTo *Wallet, amount uint64) error
	GetUserWallets(ctx context.Context, user *User) ([]*Wallet, error)
}
