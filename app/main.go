package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/bank-transaction-service/wallet/delivery/handler"
	"github.com/dabarov/bank-transaction-service/wallet/delivery/middleware"
	"github.com/dabarov/bank-transaction-service/wallet/repository/postgresql"
	"github.com/dabarov/bank-transaction-service/wallet/repository/redis"
	"github.com/dabarov/bank-transaction-service/wallet/usecase"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbusername := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhostname := os.Getenv("POSTGRES_HOSTNAME")
	dbport := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")

	redishost := os.Getenv("REDIS_HOST")
	redisport := os.Getenv("REDIS_PORT")
	redispassword := os.Getenv("REDIS_PASSWORD")
	redisdb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redissecret := os.Getenv("REDIS_SECRET")
	redistimeout, _ := strconv.Atoi(os.Getenv("REDIS_TOKEN_TIMEOUT"))

	httpport := os.Getenv("HTTP_PORT")

	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbusername, dbpassword, dbhostname, dbport, dbname)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	client := go_redis.NewClient(&go_redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redishost, redisport),
		Password: redispassword,
		DB:       redisdb,
	})

	dbRepository := postgresql.NewWalletPostgresqlRepository(db)
	redisRepository := redis.NewWalletRedisRepository(client, redistimeout, redissecret)
	walletUsecase := usecase.NewWalletUsecase(dbRepository, redisRepository)
	router := fasthttprouter.New()
	handler.NewWalletHandler(router, walletUsecase)
	log.Fatal(fasthttp.ListenAndServe(":"+httpport, middleware.NewWalletAuthMiddleware(walletUsecase, router.Handler)))
}
