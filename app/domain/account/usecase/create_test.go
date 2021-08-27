package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountCreate(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, nil)
	accountUseCase := NewAccountUseCase(&memoryStorage, zerolog.Logger{})
	fakeAccount := account.CreateAccountInput{
		Name:   "Percy Jackson",
		CPF:    "12345678913",
		Secret: "foobar",
	}
	_, _ = accountUseCase.Create(fakeAccount)

	tests := []struct {
		name       string
		inputs     account.CreateAccountInput
		wantErr    error
		wantResult entities.Account
	}{
		{
			name: "should creates an account successfully",
			inputs: account.CreateAccountInput{
				Name:   "Grover Underwood",
				CPF:    "12345678910",
				Secret: "foobar",
			},
			wantErr: nil,
			wantResult: entities.Account{
				Name:    "Grover Underwood",
				CPF:     "12345678910",
				Balance: 100,
			},
		},
		{
			name: "should return an error because name is invalid",
			inputs: account.CreateAccountInput{
				Name:   "",
				CPF:    "12345678911",
				Secret: "foobar",
			},
			wantErr:    errors.New("invalid name"),
			wantResult: entities.Account{},
		},
		{
			name: "should return an error because cpf is invalid",
			inputs: account.CreateAccountInput{
				Name:   "Grover Underwood",
				CPF:    "123456789",
				Secret: "foobar",
			},
			wantErr:    errors.New("invalid cpf"),
			wantResult: entities.Account{},
		},
		{
			name: "should return an error because secret is invalid",
			inputs: account.CreateAccountInput{
				Name:   "Grover Underwood",
				CPF:    "12345678912",
				Secret: "",
			},
			wantErr:    errors.New("invalid secret"),
			wantResult: entities.Account{},
		},
		{
			name: "should return an error because account already exist",
			inputs: account.CreateAccountInput{
				Name:   "Percy Jackson",
				CPF:    "12345678913",
				Secret: "foobar",
			},
			wantErr:    errors.New("account already exist"),
			wantResult: entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountCreated, err := accountUseCase.Create(tt.inputs)

			assert.Equal(t, tt.wantResult.Name, accountCreated.Name)
			assert.Equal(t, tt.wantResult.CPF, accountCreated.CPF)
			assert.Equal(t, tt.wantResult.Balance, accountCreated.Balance)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
			}

			if tt.wantErr == nil {
				assert.NotEqual(t, tt.inputs.Secret, accountCreated.Secret)
			}

		})
	}

}
