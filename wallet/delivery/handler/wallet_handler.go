package handler

import (
	"encoding/binary"
	"fmt"
	"strconv"

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
	router.POST("/wallet/transfer", handler.Transfer)
	router.GET("/wallet", handler.GetUserWallets)
}

func (w *WalletHandler) Create(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("userIIN"))
	if err := w.walletUsecase.Create(ctx, iin); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
}

func (w *WalletHandler) Deposit(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("userIIN"))
	amount, amountErr := strconv.ParseUint(string(ctx.FormValue("amount")), 10, 64)
	if amountErr != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", amountErr)
		return
	}
	walletID, uuidErr := uuid.ParseBytes(ctx.FormValue("walletID"))
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
	iin := fmt.Sprintf("%s", ctx.UserValue("userIIN"))
	amount := binary.BigEndian.Uint64(ctx.FormValue("amount"))
	walletFromID, uuidFromErr := uuid.ParseBytes(ctx.FormValue("walletFromID"))
	if uuidFromErr != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", uuidFromErr)
		return
	}
	walletToID, uuidToErr := uuid.ParseBytes(ctx.FormValue("walletToID"))
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

func (w *WalletHandler) GetUserWallets(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("userIIN"))
	wallets, err := w.walletUsecase.GetUserWallets(ctx, iin)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(wallets)
}
