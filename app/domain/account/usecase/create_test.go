package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountCreate(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage, authenticationStorage)
	accountUsecase := NewAccountUsecase(&memoryStorage)
	fakeAccount := CreateAccountInput{
		Name:   "Pitágoras",
		CPF:    "12345678913",
		Secret: "foobar",
	}
	accountUsecase.Create(fakeAccount)

	tests := []struct {
		name       string
		inputs     CreateAccountInput
		wantErr    error
		wantResult entities.Account
	}{
		{
			name: "should creates an account successfully",
			inputs: CreateAccountInput{
				Name:   "Tales de Mileto",
				CPF:    "12345678910",
				Secret: "foobar",
			},
			wantErr: nil,
			wantResult: entities.Account{
				Name:    "Tales de Mileto",
				CPF:     "12345678910",
				Balance: 100,
			},
		},
		{
			name: "should return an error because name is invalid",
			inputs: CreateAccountInput{
				Name:   "",
				CPF:    "12345678911",
				Secret: "foobar",
			},
			wantErr:    errors.New("invalid name"),
			wantResult: entities.Account{},
		},
		{
			name: "should return an error because cpf is invalid",
			inputs: CreateAccountInput{
				Name:   "Tales de Mileto",
				CPF:    "123456789",
				Secret: "foobar",
			},
			wantErr:    errors.New("invalid cpf"),
			wantResult: entities.Account{},
		},
		{
			name: "should return an error because secret is invalid",
			inputs: CreateAccountInput{
				Name:   "Tales de Mileto",
				CPF:    "12345678912",
				Secret: "",
			},
			wantErr:    errors.New("invalid secret"),
			wantResult: entities.Account{},
		},
		{
			name: "should return an error because account already exist",
			inputs: CreateAccountInput{
				Name:   "Pitágoras",
				CPF:    "12345678913",
				Secret: "foobar",
			},
			wantErr:    errors.New("account already exist"),
			wantResult: entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountCreated, err := accountUsecase.Create(tt.inputs)

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
