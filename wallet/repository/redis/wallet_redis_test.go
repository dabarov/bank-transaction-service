package redis_test

import (
	"os"
	"testing"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/dabarov/bank-transaction-service/wallet/repository/redis"
	goredis "github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

var (
	r        domain.WalletRedisRepository
	db       *goredis.Client
	mock     redismock.ClientMock
	secret   = "secret"
	exp_time = 10
)

func TestMain(m *testing.M) {
	db, mock = redismock.NewClientMock()
	r = redis.NewWalletRedisRepository(db, exp_time, secret)
	os.Exit(m.Run())
}
func TestGetValue(t *testing.T) {
	key := "key"
	val := "val"
	mock.ExpectGet(key).SetVal(val)
	if value, err := r.GetValue(key); err != nil || value != val {
		t.Fatalf("got error on get value: %s", err)
	}
}

func TestGetSecret(t *testing.T) {
	if value := r.GetSecret(); value != secret {
		t.Fatalf("did not get proper secret")
	}
}
