package postgresql

import (
	"context"
	"math/rand"
	"time"

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
	wallet := &domain.Wallet{
		ID:                  rand.Uint64(),
		UserIIN:             iin,
		LastTransactionDate: time.Time{}.String(),
		Balance:             0,
		CreatedAt:           time.Now().String(),
	}

	if err := w.Conn.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}

func (w *walletPostgresqlRepository) Deposit(ctx context.Context, walletID uint64, amount uint64) error {
	return nil
}

func (w *walletPostgresqlRepository) Transfer(ctx context.Context, walletFromID uint64, walletToID uint64, amount uint64) error {
	return nil
}

func (w *walletPostgresqlRepository) GetUserWallets(ctx context.Context, iin string) ([]*domain.Wallet, error) {
	return nil, nil
}
