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
	var wallet domain.Wallet
	if err := w.Conn.Where(&domain.Wallet{ID: walletID}).First(&wallet).Error; err != nil {
		return err
	}

	wallet.Balance += amount
	if err := w.Conn.Save(wallet).Error; err != nil {
		return err
	}

	return nil
}

func (w *walletPostgresqlRepository) Transfer(ctx context.Context, walletFromID uint64, walletToID uint64, amount uint64) error {
	var walletFrom domain.Wallet
	var walletTo domain.Wallet
	if err := w.Conn.Where(&domain.Wallet{ID: walletFromID}).First(&walletFrom).Error; err != nil {
		return err
	}
	if err := w.Conn.Where(&domain.Wallet{ID: walletToID}).First(&walletTo).Error; err != nil {
		return err
	}

	walletFrom.Balance -= amount
	walletTo.Balance += amount
	if err := w.Conn.Save(walletFrom).Error; err != nil {
		return err
	}
	if err := w.Conn.Save(walletTo).Error; err != nil {
		return err
	}

	return nil
}

func (w *walletPostgresqlRepository) GetUserWallets(ctx context.Context, iin string) ([]*domain.Wallet, error) {
	var wallets []*domain.Wallet
	if err := w.Conn.Where(&domain.Wallet{UserIIN: iin}).Find(&wallets).Error; err != nil {
		return wallets, err
	}
	return wallets, nil
}
