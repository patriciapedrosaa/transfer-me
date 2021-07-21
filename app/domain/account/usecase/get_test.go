package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage)
	accountUsecase := NewAccountUsecase(&memoryStorage)

	fakeAccount1 := CreateAccountInput{
		Name:   "Tales Mileto",
		CPF:    "12345678910",
		Secret: "foobar",
	}
	fakeAccount2 := CreateAccountInput{
		Name:   "Pitágoras",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	_, _ = accountUsecase.Create(fakeAccount1)
	_, _ = accountUsecase.Create(fakeAccount2)

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
					Name:    "Tales Mileto",
					CPF:     "12345678910",
					Balance: 100,
				},
				{
					Name:    "Pitágoras",
					CPF:     "12345678911",
					Balance: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountUsecase.GetAccounts()

			for k, _ := range got {
				assert.Equal(t, tt.wantResult[k].Name, got[k].Name)
				assert.Equal(t, tt.wantResult[k].Name, got[k].Name)
				assert.Equal(t, tt.wantResult[k].CPF, got[k].CPF)
				assert.Equal(t, tt.wantResult[k].Balance, got[k].Balance)
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
