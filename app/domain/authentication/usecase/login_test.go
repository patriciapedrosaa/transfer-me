package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckLogin(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, authenticationStorage)
	authenticationUsecase := NewAuthenticationUseCase(&memoryStorage, &memoryStorage)
	accountUsecase := au.NewAccountUseCase(&memoryStorage)

	accountTest := account.CreateAccountInput{
		Name:   "Isaac Newton",
		CPF:    "12345678910",
		Secret: "foobar",
	}
	accountCreated, _ := accountUsecase.Create(accountTest)

	t.Run("should return valid login successfully", func(t *testing.T) {
		inputs := LoginInputs{
			CPF:     "12345678910",
			Secret:  "foobar",
			Account: accountCreated,
		}
		got, err := authenticationUsecase.CheckLogin(inputs)
		assert.Nil(t, err)
		assert.True(t, got)
	})

	tests := []struct {
		name      string
		inputs    LoginInputs
		wantError error
	}{
		{
			name: "should return an error because CPF is invalid",
			inputs: LoginInputs{
				CPF:    "12345678911",
				Secret: "foobar",
			},
			wantError: ErrInvalidCredentials,
		},
		{
			name: "should return an error because secret is invalid",
			inputs: LoginInputs{
				CPF:    "12345678910",
				Secret: "foo",
			},
			wantError: ErrInvalidCredentials,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authenticationUsecase.CheckLogin(tt.inputs)

			assert.Equal(t, tt.wantError, err)
			assert.False(t, got)

		})
	}

}
