package postgresql

import (
	"context"
	"time"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type walletPostgresqlRepository struct {
	Conn *gorm.DB
}

func NewWalletPostgresqlRepository(Conn *gorm.DB) domain.WalletDBRepository {
	Conn.AutoMigrate(&domain.Wallet{})
	return &walletPostgresqlRepository{Conn}
}

func (w *walletPostgresqlRepository) Create(ctx context.Context, iin string) error {
	wallet := &domain.Wallet{
		UserIIN:             iin,
		LastTransactionDate: time.Time{}.String(),
		Balance:             0,
		CreatedAt:           time.Now().String(),
	}

	err := w.Conn.Create(&wallet).Error
	return err
}

func (w *walletPostgresqlRepository) Deposit(ctx context.Context, iin string, walletID uuid.UUID, amount uint64) error {
	var wallet *domain.Wallet

	if err := w.Conn.Where(&domain.Wallet{ID: walletID}).First(&wallet).Error; err != nil {
		return err
	}

	err := w.Conn.Model(&wallet).Update("balance", wallet.Balance+amount).Error
	return err
}
