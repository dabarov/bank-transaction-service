package postgresql

import (
	"context"

	"github.com/dabarov/bank-transaction-service/domain"
	"gorm.io/gorm"
)

type walletPostgresqlRepository struct {
	Conn *gorm.DB
}

func NewWalletPostgresqlRepository(Conn *gorm.DB) domain.WalletRepository {
	return &walletPostgresqlRepository{Conn}
}

func (w *walletPostgresqlRepository) Create(ctx context.Context, iin string) error {

	return nil
}

func (w *walletPostgresqlRepository) Deposit(ctx context.Context, wallet *domain.Wallet, amount uint64) error {
	return nil
}

func (w *walletPostgresqlRepository) Transfer(ctx context.Context, walletFrom *domain.Wallet, walletTo *domain.Wallet, amount uint64) error {
	return nil
}

func (w *walletPostgresqlRepository) GetUserWallets(ctx context.Context, user *domain.User) ([]*domain.Wallet, error) {
	return nil, nil
}
