package http

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/valyala/fasthttp"
)

type WalletHandler struct {
	WUsecase domain.WalletUsecase
}

func NewWalletHandler(router *fasthttprouter.Router, wcase domain.WalletUsecase) {
	handler := &WalletHandler{
		WUsecase: wcase,
	}
	router.POST("/wallet/create", handler.Create)
	router.PUT("/wallet/deposit", handler.Deposit)
	router.PUT("/wallet/transfer", handler.Transfer)
	router.GET("/wallet/", handler.GetUserWallets)
}

func (w *WalletHandler) Create(ctx *fasthttp.RequestCtx) {
}

func (w *WalletHandler) Deposit(ctx *fasthttp.RequestCtx) {
}

func (w *WalletHandler) Transfer(ctx *fasthttp.RequestCtx) {
}

func (w *WalletHandler) GetUserWallets(ctx *fasthttp.RequestCtx) {
}
