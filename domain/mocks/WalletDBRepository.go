package mocks

import (
	"context"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type WalletDBRepository struct {
	mock.Mock
}

func (_m *WalletDBRepository) Create(ctx context.Context, iin string) error {
	ret := _m.Called(ctx, iin)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, iin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *WalletDBRepository) Deposit(ctx context.Context, iin string, walletID uuid.UUID, amount uint64) error {
	ret := _m.Called(ctx, iin, walletID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID, uint64) error); ok {
		r0 = rf(ctx, iin, walletID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *WalletDBRepository) Transfer(ctx context.Context, iin string, walletFromID uuid.UUID, walletToID uuid.UUID, amount uint64) error {
	ret := _m.Called(ctx, iin, walletFromID, walletToID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID, uuid.UUID, uint64) error); ok {
		r0 = rf(ctx, iin, walletFromID, walletToID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *WalletDBRepository) GetUserWallets(ctx context.Context, iin string) ([]domain.Wallet, error) {
	ret := _m.Called(ctx, iin)

	var r0 []domain.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Wallet); ok {
		r0 = rf(ctx, iin)
	} else {
		r0 = ret.Get(0).([]domain.Wallet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, iin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
