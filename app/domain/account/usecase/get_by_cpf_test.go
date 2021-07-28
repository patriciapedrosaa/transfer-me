package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetById(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, nil)
	accountUsecase := NewAccountUsecase(&memoryStorage)

	fakeAccount1 := CreateAccountInput{
		Name:   "Tales Mileto",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	fakeAccount2 := CreateAccountInput{
		Name:   "Pitágoras",
		CPF:    "12345678910",
		Secret: "foobar",
	}
	_, _ = accountUsecase.Create(fakeAccount1)
	_, _ = accountUsecase.Create(fakeAccount2)

	tests := []struct {
		name       string
		cpf        string
		wantErr    error
		wantResult entities.Account
	}{
		{
			name:    "should return an account successfully",
			cpf:     "12345678911",
			wantErr: nil,
			wantResult: entities.Account{
				Name:    "Tales Mileto",
				CPF:     "12345678911",
				Balance: 100,
			},
		},
		{
			name:       "should return err not found",
			cpf:        "12345678912",
			wantErr:    errors.New("not found"),
			wantResult: entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountUsecase.GetByCpf(tt.cpf)

			assert.Equal(t, tt.wantResult.Name, got.Name)
			assert.Equal(t, tt.wantResult.CPF, got.CPF)
			assert.Equal(t, tt.wantResult.Balance, got.Balance)
			assert.Equal(t, tt.wantErr, err)
			//assert.Error(t, tt.wantErr, err)
		})
	}
}
