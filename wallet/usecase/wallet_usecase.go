package usecase

import (
	"context"
	"time"

	"github.com/dabarov/bank-transaction-service/domain"
)

type walletUsecase struct {
	walletRepository *domain.WalletRepository
	contextTimeout   time.Duration
}

func NewWalletUsecase(w *domain.WalletRepository, timeout time.Duration) domain.WalletUsecase {
	return &walletUsecase{
		walletRepository: w,
		contextTimeout:   timeout,
	}
}

func (w *walletUsecase) Create(ctx context.Context, iin uint64) error {
	return nil
}

func (w *walletUsecase) Deposit(ctx context.Context, wallet *domain.Wallet, amount uint64) error {
	return nil
}

func (w *walletUsecase) Transfer(ctx context.Context, walletFrom *domain.Wallet, walletTo *domain.Wallet, amount uint64) error {
	return nil
}

func (w *walletUsecase) GetUserWallets(ctx context.Context, user *domain.User) ([]*domain.Wallet, error) {
	return nil, nil
}
