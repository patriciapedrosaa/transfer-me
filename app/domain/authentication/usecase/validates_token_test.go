package usecase

import (
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatesToken(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage, authenticationStorage)
	accountUsecase := au.NewAccountUsecase(&memoryStorage)
	authenticationUsecase := NewTransferUsecase(&memoryStorage, &memoryStorage)

	accountTest := au.CreateAccountInput{
		Name:   "Patricia",
		CPF:    "12345678918",
		Secret: "foobar",
	}
	_, _ = accountUsecase.Create(accountTest)
	user1 := LoginInputs{
		CPF:    "12345678918",
		Secret: "foobar",
	}

	token, _ := authenticationUsecase.CreateToken(user1)

	t.Run("Should return a token successfully", func(t *testing.T) {
		got, err := authenticationUsecase.ValidatesToken(token)

		assert.NotEmpty(t, got)
		assert.Empty(t, err)
	})

	t.Run("Should return error because token is invalid", func(t *testing.T) {
		wrongToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		got, err := authenticationUsecase.ValidatesToken(wrongToken)

		assert.Empty(t, got)
		assert.Error(t, err)
	})
}
