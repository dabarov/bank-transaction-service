package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/dabarov/bank-transaction-service/domain/mocks"
	"github.com/dabarov/bank-transaction-service/wallet/usecase"
	"github.com/google/uuid"
)

var (
	dbRepo    = new(mocks.WalletDBRepository)
	redisRepo = new(mocks.WalletRedisRepository)
	mockCtx   = context.Background()
)

func TestGetRedisSecret(t *testing.T) {
	redisRepo.On("GetSecret").Return("secret").Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if secret := mockUseCase.GetRedisSecret(); secret != "secret" {
		t.Fatalf("should have gotten secret")
	}
}

func TestGetRedisValue(t *testing.T) {
	key := "key"
	value := "val"
	redisRepo.On("GetValue", key).Return(value, nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if valCheck, err := mockUseCase.GetRedisValue(key); valCheck != value || err != nil {
		t.Fatalf("should have gotten value")
	}
}

func TestCreate(t *testing.T) {
	iin := "123123123123"
	dbRepo.On("Create", mockCtx, iin).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Create(mockCtx, iin); err != nil {
		t.Fatalf("got error on create %v", err)
	}
}

func TestCreateWrongIIN(t *testing.T) {
	iin := "12s3123123123"
	dbRepo.On("Create", mockCtx, iin).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Create(mockCtx, iin); err != domain.ErrIINIncorect {
		t.Fatalf("should have received iin incorrect error")
	}
}

func TestDeposit(t *testing.T) {
	iin := "123123123123"
	walletID, _ := uuid.NewUUID()
	var amount uint64 = 100
	dbRepo.On("Deposit", mockCtx, iin, walletID, amount).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Deposit(mockCtx, iin, walletID, amount); err != nil {
		t.Fatalf("got error during deposit %v", err)
	}
}

func TestDepositWrongIIN(t *testing.T) {
	iin := "123"
	walletID, _ := uuid.NewUUID()
	var amount uint64 = 100
	dbRepo.On("Deposit", mockCtx, iin, walletID, amount).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Deposit(mockCtx, iin, walletID, amount); err != domain.ErrIINIncorect {
		t.Fatalf("missed test on incorrect iin")
	}
}

func TestDepositZeroAmount(t *testing.T) {
	iin := "123123123123"
	walletID, _ := uuid.NewUUID()
	var amount uint64 = 0
	dbRepo.On("Deposit", mockCtx, iin, walletID, amount).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Deposit(mockCtx, iin, walletID, amount); err != domain.ErrZeroAmount {
		t.Fatalf("missed test on zero amount")
	}
}

func TestTransfer(t *testing.T) {
	iin := "123123123123"
	walletFromID, _ := uuid.NewUUID()
	walletToID, _ := uuid.NewUUID()
	var amount uint64 = 100
	dbRepo.On("Transfer", mockCtx, iin, walletFromID, walletToID, amount).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Transfer(mockCtx, iin, walletFromID, walletToID, amount); err != nil {
		t.Fatalf("got error during transfer %v", err)
	}
}

func TestTransferWrongIIN(t *testing.T) {
	iin := "12312312312311"
	walletFromID, _ := uuid.NewUUID()
	walletToID, _ := uuid.NewUUID()
	var amount uint64 = 100
	dbRepo.On("Transfer", mockCtx, iin, walletFromID, walletToID, amount).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Transfer(mockCtx, iin, walletFromID, walletToID, amount); err != domain.ErrIINIncorect {
		t.Fatalf("missed incorrect iin test")
	}
}

func TestTransferZeroAmount(t *testing.T) {
	iin := "123123123123"
	walletFromID, _ := uuid.NewUUID()
	walletToID, _ := uuid.NewUUID()
	var amount uint64 = 0
	dbRepo.On("Transfer", mockCtx, iin, walletFromID, walletToID, amount).Return(nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if err := mockUseCase.Transfer(mockCtx, iin, walletFromID, walletToID, amount); err != domain.ErrZeroAmount {
		t.Fatalf("missed zero amount test")
	}
}

func TestGetUserWallets(t *testing.T) {
	iin := "123123123123"
	dbRepo.On("GetUserWallets", mockCtx, iin).Return([]domain.Wallet{}, nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if _, err := mockUseCase.GetUserWallets(mockCtx, iin); err != nil {
		t.Fatalf("got error during get user wallet %v", err)
	}
}

func TestGetUserWalletsWrongIIN(t *testing.T) {
	iin := "123123123123111"
	dbRepo.On("GetUserWallets", mockCtx, iin).Return([]domain.Wallet{}, nil).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if _, err := mockUseCase.GetUserWallets(mockCtx, iin); err != domain.ErrIINIncorect {
		t.Fatalf("missed incorrect iin test")
	}
}

func TestGetUserWalletsDBError(t *testing.T) {
	iin := "123123123123"
	dbError := errors.New("db Error")
	dbRepo.On("GetUserWallets", mockCtx, iin).Return([]domain.Wallet{}, dbError).Once()
	mockUseCase := usecase.NewWalletUsecase(dbRepo, redisRepo)
	if _, err := mockUseCase.GetUserWallets(mockCtx, iin); err != dbError {
		t.Fatalf("missed db error test")
	}
}
