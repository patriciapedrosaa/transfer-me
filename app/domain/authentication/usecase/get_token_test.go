package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage, authenticationStorage)
	authenticationUsecase := NewAuthenticationUseCase(&memoryStorage, &memoryStorage)

	token := memory.Token{
		ID:        uuid.New().String(),
		Subject:   uuid.New().String(),
		Issuer:    "JWT",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 15),
	}

	authenticationStorage[token.ID] = token

	t.Run("Should return a token successfully", func(t *testing.T) {
		got, err := authenticationUsecase.GetToken(token.ID)

		assert.Equal(t, got.Subject, token.Subject)
		assert.Equal(t, got.Issuer, token.Issuer)
		assert.Equal(t, got.Name, token.Name)
		assert.Empty(t, err)
	})

	t.Run("should return error not found", func(t *testing.T) {
		fakeId := uuid.New().String()
		got, err := authenticationUsecase.GetToken(fakeId)

		assert.Error(t, err)
		assert.Empty(t, got)
		assert.Equal(t, err, errors.New("not found"))

	})
}
