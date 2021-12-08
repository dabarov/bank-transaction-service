package usecase

import (
	"context"

	"github.com/dabarov/bank-transaction-service/domain"
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

func (w *walletUsecase) GetRedisValue(key string) (string, error) {
	return w.walletRedisRepository.GetValue(key)
}

func (w *walletUsecase) GetRedisSecret() string {
	return w.walletRedisRepository.GetSecret()
}
