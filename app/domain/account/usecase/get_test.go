package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, nil)
	accountUseCase := NewAccountUseCase(&memoryStorage, zerolog.Logger{})

	fakeAccount1 := account.CreateAccountInput{
		Name:   "Ronald Weasley",
		CPF:    "12345678910",
		Secret: "foobar",
	}
	fakeAccount2 := account.CreateAccountInput{
		Name:   "Ginevra Weasley",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	_, _ = accountUseCase.Create(fakeAccount1)
	_, _ = accountUseCase.Create(fakeAccount2)

	tests := []struct {
		name       string
		wantErr    error
		wantResult []entities.Account
	}{
		{
			name:    "should return a list of account",
			wantErr: nil,
			wantResult: []entities.Account{
				{
					Name:    "Ronald Weasley",
					CPF:     "12345678910",
					Balance: 100,
				},
				{
					Name:    "Ginevra Weasley",
					CPF:     "12345678911",
					Balance: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountUseCase.GetAccounts()

			for k := range got {
				assert.Equal(t, tt.wantResult[k].Name, got[k].Name)
				assert.Equal(t, tt.wantResult[k].Name, got[k].Name)
				assert.Equal(t, tt.wantResult[k].CPF, got[k].CPF)
				assert.Equal(t, tt.wantResult[k].Balance, got[k].Balance)
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
