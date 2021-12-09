package usecase

import (
	"context"
	"encoding/json"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/google/uuid"
)

type walletUsecase struct {
	walletRepository      domain.WalletDBRepository
	walletRedisRepository domain.WalletRedisRepository
}

func NewWalletUsecase(wDBR domain.WalletDBRepository, wRR domain.WalletRedisRepository) domain.WalletUsecase {
	return &walletUsecase{
		walletRepository:      wDBR,
		walletRedisRepository: wRR,
	}
}

func (w *walletUsecase) Create(ctx context.Context, iin string) error {
	if InvalidIIN(iin) {
		return domain.ErrIINIncorect
	}
	err := w.walletRepository.Create(ctx, iin)
	return err
}

func (w *walletUsecase) Deposit(ctx context.Context, iin string, walletID uuid.UUID, amount uint64) error {
	if InvalidIIN(iin) {
		return domain.ErrIINIncorect
	}
	if ZeroAmount(amount) {
		return domain.ErrZeroAmount
	}
	err := w.walletRepository.Deposit(ctx, iin, walletID, amount)
	return err
}

func (w *walletUsecase) Transfer(ctx context.Context, iin string, walletFromID uuid.UUID, walletToID uuid.UUID, amount uint64) error {
	if InvalidIIN(iin) {
		return domain.ErrIINIncorect
	}
	if ZeroAmount(amount) {
		return domain.ErrZeroAmount
	}
	err := w.walletRepository.Transfer(ctx, iin, walletFromID, walletToID, amount)
	return err
}

func (w *walletUsecase) GetUserWallets(ctx context.Context, iin string) ([]byte, error) {
	if InvalidIIN(iin) {
		return nil, domain.ErrIINIncorect
	}
	walletArray, err := w.walletRepository.GetUserWallets(ctx, iin)
	if err != nil {
		return nil, err
	}
	responseJSON, jsonErr := json.Marshal(walletArray)
	return responseJSON, jsonErr
}

func (w *walletUsecase) GetRedisValue(key string) (string, error) {
	return w.walletRedisRepository.GetValue(key)
}

func (w *walletUsecase) GetRedisSecret() string {
	return w.walletRedisRepository.GetSecret()
}
