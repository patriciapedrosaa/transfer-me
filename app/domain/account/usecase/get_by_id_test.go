package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetById(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, nil)
	accountUseCase := NewAccountUseCase(&memoryStorage, zerolog.Logger{})

	fakeAccount1 := account.CreateAccountInput{
		Name:   "Dino Thomas",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	account1, _ := accountUseCase.Create(fakeAccount1)

	tests := []struct {
		name       string
		id         string
		wantErr    error
		wantResult entities.Account
	}{
		{
			name:    "should return an account successfully",
			id:      account1.AccountID,
			wantErr: nil,
			wantResult: entities.Account{
				Name:    "Dino Thomas",
				CPF:     "12345678911",
				Balance: 100,
			},
		},
		{
			name:       "should return err not found",
			id:         uuid.New().String(),
			wantErr:    errors.New("not found"),
			wantResult: entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountUseCase.GetById(tt.id)

			assert.Equal(t, tt.wantResult.Name, got.Name)
			assert.Equal(t, tt.wantResult.CPF, got.CPF)
			assert.Equal(t, tt.wantResult.Balance, got.Balance)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
