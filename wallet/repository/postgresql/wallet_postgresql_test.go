package postgresql_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/dabarov/bank-transaction-service/wallet/repository/postgresql"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db      *sql.DB
	mock    sqlmock.Sqlmock
	repo    domain.WalletDBRepository
	testCtx = context.Background()
)

func TestMain(m *testing.M) {
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})
	repo = postgresql.NewWalletPostgresqlRepository(gormDB)
	if err != nil {
		log.Fatalf("%v", err)
	}
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	iin := "iin"

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "wallets" ("id","user_iin","last_transaction_date","balance","created_at") VALUES ($1,$2,$3,$4,$5)`)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := repo.Create(testCtx, iin); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestCreateFail(t *testing.T) {
	iin := "iin"

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectCommit()

	if err := repo.Create(testCtx, iin); err == nil {
		t.Fatalf("%v", err)
	}
}

func TestDeposit(t *testing.T) {
	iin := "iin"
	uid, err := uuid.NewUUID()
	var amount uint64 = 100
	if err != nil {
		t.Fatalf("%v", err)
	}
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(uid).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(uid, iin, "date", amount, "date"))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "wallets" SET "balance"=$1 WHERE "id" = $2`)).WithArgs(200, uid).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := repo.Deposit(testCtx, iin, uid, amount); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestDepositFail(t *testing.T) {
	iin := "iin"
	uid, err := uuid.NewUUID()
	var amount uint64 = 100
	if err != nil {
		t.Fatalf("%v", err)
	}
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(uid)
	mock.ExpectCommit()

	if err := repo.Deposit(testCtx, iin, uid, amount); err == nil {
		t.Fatalf("%v", err)
	}
}

func TestGetUserWallet(t *testing.T) {
	iin := "iin"
	uid, err := uuid.NewUUID()
	var amount uint64 = 100
	if err != nil {
		t.Fatalf("%v", err)
	}
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(iin).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(uid, iin, "date", amount, "date"))
	mock.ExpectCommit()

	if _, err := repo.GetUserWallets(testCtx, iin); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetUserWalletFail(t *testing.T) {
	iin := "iin"
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(iin)
	mock.ExpectCommit()

	if _, err := repo.GetUserWallets(testCtx, iin); err == nil {
		t.Fatalf("%v", err)
	}
}

func TestTransfer(t *testing.T) {
	iin := ""
	walletFromID, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	walletToID, _ := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	var amount uint64 = 100

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(walletFromID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletFromID, iin, "date", amount, "date"))
	mock.ExpectQuery("SELECT").WithArgs(walletToID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletToID, iin, "date", amount, "date"))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "wallets" SET "balance"=$1 WHERE "id" = $2`)).WithArgs(0, walletFromID).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "wallets" SET "balance"=$1 WHERE "id" = $2`)).WithArgs(200, walletToID).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	if err := repo.Transfer(testCtx, iin, walletFromID, walletToID, amount); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestTransferNotFoundFrom(t *testing.T) {
	iin := ""
	walletFromID, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	walletToID, _ := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	var amount uint64 = 100

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(walletFromID)
	mock.ExpectCommit()
	if err := repo.Transfer(testCtx, iin, walletFromID, walletToID, amount); err == nil {
		t.Fatalf("should have not found from wallet")
	}
}

func TestTransferNotFoundTo(t *testing.T) {
	iin := ""
	walletFromID, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	walletToID, _ := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	var amount uint64 = 100

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(walletFromID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletFromID, iin, "date", amount, "date"))
	mock.ExpectQuery("SELECT").WithArgs(walletToID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}))
	mock.ExpectCommit()
	if err := repo.Transfer(testCtx, iin, walletFromID, walletToID, amount); err == nil {
		t.Fatalf("should have not found wallet to")
	}
}

func TestTransferNotMatchinIIN(t *testing.T) {
	iin := "someIIN"
	walletFromID, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	walletToID, _ := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	var amount uint64 = 100

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(walletFromID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletFromID, "anotherIIN", "date", amount, "date"))
	mock.ExpectQuery("SELECT").WithArgs(walletToID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletFromID, "anotherIIN", "date", amount, "date"))
	mock.ExpectCommit()
	if err := repo.Transfer(testCtx, iin, walletFromID, walletToID, amount); err == nil {
		t.Fatalf("should have not match iins")
	}
}

func TestTransferNotUpdateFrom(t *testing.T) {
	iin := ""
	walletFromID, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	walletToID, _ := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	var amount uint64 = 100

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(walletFromID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletFromID, iin, "date", amount, "date"))
	mock.ExpectQuery("SELECT").WithArgs(walletToID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletToID, iin, "date", amount, "date"))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "wallets" SET "balance"=$1 WHERE "id" = $2`)).WithArgs(0, walletFromID)
	mock.ExpectCommit()
	if err := repo.Transfer(testCtx, iin, walletFromID, walletToID, amount); err == nil {
		t.Fatalf("should have not updated from")
	}
}

func TestTransferNotUpdateTo(t *testing.T) {
	iin := ""
	walletFromID, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	walletToID, _ := uuid.NewUUID()
	if err != nil {
		t.Fatalf("%v", err)
	}
	var amount uint64 = 100

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(walletFromID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletFromID, iin, "date", amount, "date"))
	mock.ExpectQuery("SELECT").WithArgs(walletToID).WillReturnRows(sqlmock.NewRows([]string{"id", "user", "last_transaction_date", "balance", "created_at"}).AddRow(walletToID, iin, "date", amount, "date"))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "wallets" SET "balance"=$1 WHERE "id" = $2`)).WithArgs(0, walletFromID).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "wallets" SET "balance"=$1 WHERE "id" = $2`)).WithArgs(200, walletToID)
	mock.ExpectCommit()
	if err := repo.Transfer(testCtx, iin, walletFromID, walletToID, amount); err == nil {
		t.Fatalf("should have not updated wallet to")
	}
}
