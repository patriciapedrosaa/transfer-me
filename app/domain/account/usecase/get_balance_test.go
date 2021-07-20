package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBlance(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage)
	accountUsecase := NewAccountUsecase(&memoryStorage)
	fakeAccount1 := CreateAccountInput{
		Name:   "Tales Mileto",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	_, _ = accountUsecase.Create(fakeAccount1)

	tests := []struct {
		name       string
		cpf        string
		wantErr    error
		wantResult int
	}{
		{
			name:       "should return balance successfully",
			cpf:        "12345678911",
			wantErr:    nil,
			wantResult: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountUsecase.GetBalance(tt.cpf)

			assert.Equal(t, tt.wantResult, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
