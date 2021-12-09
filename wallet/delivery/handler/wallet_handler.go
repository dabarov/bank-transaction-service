package handler

import (
	"encoding/binary"
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/google/uuid"
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
	router.POST("/wallet/deposit", handler.Deposit)
}

func (w *WalletHandler) Create(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("iin"))
	if err := w.walletUsecase.Create(ctx, iin); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
}

func (w *WalletHandler) Deposit(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("iin"))
	amount := binary.BigEndian.Uint64(ctx.FormValue("amount"))
	walletID, uuidErr := uuid.FromBytes(ctx.FormValue("walletID"))
	if uuidErr != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", uuidErr)
		return
	}
	if err := w.walletUsecase.Deposit(ctx, iin, walletID, amount); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
}

func (w *WalletHandler) Transfer(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("iin"))
	amount := binary.BigEndian.Uint64(ctx.FormValue("amount"))
	walletFromID, uuidFromErr := uuid.FromBytes(ctx.FormValue("walletID"))
	if uuidFromErr != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", uuidFromErr)
		return
	}
	walletToID, uuidToErr := uuid.FromBytes(ctx.FormValue("walletID"))
	if uuidToErr != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", uuidToErr)
		return
	}

	if err := w.walletUsecase.Transfer(ctx, iin, walletFromID, walletToID, amount); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
}
