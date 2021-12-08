package handler

import (
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/valyala/fasthttp"
)

type WalletHandler struct {
	walletUsecase domain.WalletUsecase
}

func NewWalletHandler(router *fasthttprouter.Router, wcase domain.WalletUsecase) {
	handler := &WalletHandler{
		walletUsecase: wcase,
	}
	router.POST("/wallet/create/:iin", handler.Create)
}

func (w *WalletHandler) Create(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("iin"))
	if err := w.walletUsecase.Create(ctx, iin); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
}
