package middleware

import (
	"fmt"
	"log"

	"github.com/dabarov/bank-transaction-service/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

type walletAuthMiddleware struct {
	walletUsecase domain.WalletUsecase
}

func NewCORSMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers, X-Custom-Header")
		next(ctx)
	}
}

func NewWalletAuthMiddleware(walletUsecase domain.WalletUsecase, next fasthttp.RequestHandler) fasthttp.RequestHandler {
	middleware := &walletAuthMiddleware{
		walletUsecase: walletUsecase,
	}
	return func(ctx *fasthttp.RequestCtx) {
		token, err := middleware.ExtractToken(ctx)
		if err != nil {
			log.Printf("Extract token error: %v", err)
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		iin, err := middleware.ParseToken(token)
		if err != nil {
			log.Printf("Parse token error: %v", err)
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		ok := middleware.FindToken(token, iin)
		if !ok {
			log.Printf("Getting token failed")
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		ctx.SetUserValue("userIIN", iin)

		next(ctx)
	}
}

func (w *walletAuthMiddleware) FindToken(token string, iin string) bool {
	key := fmt.Sprintf("user:%s", iin)

	value, err := w.walletUsecase.GetRedisValue(key)
	if err != nil {
		return false
	}

	return token == value
}

func (w *walletAuthMiddleware) ExtractToken(ctx *fasthttp.RequestCtx) (token string, err error) {
	token = string(ctx.Request.Header.Cookie("AuthToken"))
	if token == "" {
		err = fmt.Errorf("authorization header not found")
		return
	}
	err = nil
	return
}

func (w *walletAuthMiddleware) ParseToken(token string) (string, error) {
	JWTToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to extract token metadata, unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(w.walletUsecase.GetRedisSecret()), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := JWTToken.Claims.(jwt.MapClaims)

	var userINN string

	if ok && JWTToken.Valid {
		userINN, ok = claims["iin"].(string)
		if !ok {
			return "", fmt.Errorf("field id not found")
		}
		return userINN, nil
	}

	return "", fmt.Errorf("invalid token")
}
