package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type WalletUsecase struct {
	mock.Mock
}

func (_m *WalletUsecase) Create(ctx context.Context, iin string) error {
	ret := _m.Called(ctx, iin)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, iin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *WalletUsecase) Deposit(ctx context.Context, iin string, walletID uuid.UUID, amount uint64) error {
	ret := _m.Called(ctx, iin, walletID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID, uint64) error); ok {
		r0 = rf(ctx, iin, walletID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *WalletUsecase) Transfer(ctx context.Context, iin string, walletFromID uuid.UUID, walletToID uuid.UUID, amount uint64) error {
	ret := _m.Called(ctx, iin, walletFromID, walletToID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID, uuid.UUID, uint64) error); ok {
		r0 = rf(ctx, iin, walletFromID, walletToID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *WalletUsecase) GetUserWallets(ctx context.Context, iin string) ([]byte, error) {
	ret := _m.Called(ctx, iin)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, iin)
	} else {
		r0 = ret.Get(0).([]byte)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, iin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *WalletUsecase) GetRedisValue(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *WalletUsecase) GetRedisSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
