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

func (w *walletPostgresqlRepository) Transfer(ctx context.Context, iin string, walletFromID uuid.UUID, walletToID uuid.UUID, amount uint64) error {
	tx := w.Conn.Begin()
	var walletFrom *domain.Wallet
	var walletTo *domain.Wallet

	if err := w.Conn.Where(&domain.Wallet{ID: walletFromID}).First(&walletFrom).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := w.Conn.Where(&domain.Wallet{ID: walletToID}).First(&walletTo).Error; err != nil {
		tx.Rollback()
		return err
	}

	if walletFrom.UserIIN != iin {
		tx.Rollback()
		return domain.ErrIINTransferDenied
	}

	if walletFrom.Balance < amount {
		tx.Rollback()
		return domain.ErrNotEnoughMoney
	}

	if err := w.Conn.Model(&walletFrom).Update("balance", walletFrom.Balance-amount).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := w.Conn.Model(&walletTo).Update("balance", walletTo.Balance+amount).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (w *walletPostgresqlRepository) GetUserWallets(ctx context.Context, iin string) ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	if notFound := w.Conn.Where(&domain.Wallet{UserIIN: iin}).Find(&wallets).Error; notFound != nil {
		return wallets, domain.ErrNoWalletsFound
	}
	return wallets, nil
}
