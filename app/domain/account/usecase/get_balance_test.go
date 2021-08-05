package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, nil)
	accountUseCase := NewAccountUseCase(&memoryStorage)
	fakeAccount1 := account.CreateAccountInput{
		Name:   "Sirius Black",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	account1, _ := accountUseCase.Create(fakeAccount1)

	tests := []struct {
		name       string
		id         string
		wantErr    error
		wantResult int
	}{
		{
			name:       "should return balance successfully",
			id:         account1.AccountID,
			wantErr:    nil,
			wantResult: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountUseCase.GetBalance(tt.id)

			assert.Equal(t, tt.wantResult, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
