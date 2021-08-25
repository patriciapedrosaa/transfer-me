package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestValidatesToken(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, authenticationStorage)
	accountUseCase := au.NewAccountUseCase(&memoryStorage)
	authenticationUseCase := NewAuthenticationUseCase(&memoryStorage)

	accountTest := account.CreateAccountInput{
		Name:   "Patricia",
		CPF:    "12345678918",
		Secret: "foobar",
	}
	accountCreated, _ := accountUseCase.Create(accountTest)
	user1 := authentication.LoginInputs{
		CPF:     "12345678918",
		Secret:  "foobar",
		Account: accountCreated,
	}

	token, _ := authenticationUseCase.CreateToken(user1)

	t.Run("Should return a token successfully", func(t *testing.T) {
		got, err := authenticationUseCase.ValidatesToken(token)

		assert.NotEmpty(t, got)
		assert.NotEmpty(t, got.ID)
		assert.NotEmpty(t, got.IssuedAt)
		assert.NotEmpty(t, got.ExpiredAt)
		assert.Equal(t, got.Subject, accountCreated.AccountID)
		assert.Equal(t, got.Name, accountCreated.Name)
		assert.Equal(t, got.Issuer, entities.ISSUER)
		assert.Empty(t, err)
	})

	t.Run("Should return error because token signature method is invalid", func(t *testing.T) {
		wrongToken := "eyJhbGciOiJQUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.VMOk4ckLBFstkXMK0FuApNcG2FqNjs0_D8YubBKDOJ09IQV5XEexJBUv9YYkf60JBphZw_puMMEYlOGzlvgTCNeVCmzCDTPG2mvyuUG80ZPM-3B_uZyt23TbHKNF5GFvDa0X3Fa-aXTrM4cwjMVSku0YEbTKNvN1Ei3tyuQaPEWFRG-0Z6X_7ATSDYjrhmOk-RKP6dj5Yd2f4xMPf1ab4u9u98HFHBubmXR0dl9HmnVPfOwGCn0DuA9YqfG_NEzDaUVFTWsoBIajDYDZSbtHFp-D5ylE3WbomkaYjxpkZAAHAXyXwExW1QM3FM_JZZhmMywuOuIa0gZJAwOUvXuoyg"
		wantErr := ErrMethodInvalid
		got, err := authenticationUseCase.ValidatesToken(wrongToken)

		assert.Empty(t, got)
		assert.Error(t, err)
		assert.Equal(t, wantErr, err)
	})

	t.Run("Should return error because token expired", func(t *testing.T) {
		expiredToken := generateExpiredToken(accountCreated.Name, accountCreated.AccountID)
		wantErr := errors.New("Token is expired")

		got, err := authenticationUseCase.ValidatesToken(expiredToken)

		assert.Empty(t, got)
		assert.Error(t, err)
		assert.Equal(t, wantErr, err)
	})

	t.Run("Should return error because token is not in database", func(t *testing.T) {
		fakeToken := generateFakeToken()
		wantErr := ErrTokenNotFound

		got, err := authenticationUseCase.ValidatesToken(fakeToken)

		assert.Empty(t, got)
		assert.Error(t, err)
		assert.Equal(t, wantErr, err)
	})
}

func generateExpiredToken(username string, subject string) string {
	token := entities.Token{
		ID:       uuid.New().String(),
		Name:     username,
		Subject:  subject,
		Issuer:   entities.ISSUER,
		IssuedAt: time.Time{},
	}
	token.ExpiredAt = token.IssuedAt.Add(time.Minute * 20)

	atClaims := jwt.MapClaims{
		"id":   token.ID,
		"name": token.Name,
		"sub":  token.Subject,
		"iss":  token.Issuer,
		"iat":  token.IssuedAt.Unix(),
		"exp":  token.ExpiredAt.Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return accessTokenString
}

func generateFakeToken() string {
	token := entities.Token{
		ID:       uuid.New().String(),
		Name:     "username",
		Subject:  uuid.New().String(),
		Issuer:   entities.ISSUER,
		IssuedAt: time.Now(),
	}
	token.ExpiredAt = token.IssuedAt.Add(time.Minute * entities.DURATION)

	atClaims := jwt.MapClaims{
		"id":   token.ID,
		"name": token.Name,
		"sub":  token.Subject,
		"iss":  token.Issuer,
		"iat":  token.IssuedAt.Unix(),
		"exp":  token.ExpiredAt.Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return accessTokenString
}
