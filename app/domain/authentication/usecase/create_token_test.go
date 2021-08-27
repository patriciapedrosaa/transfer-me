package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateToken(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, authenticationStorage)
	authenticationUseCase := NewAuthenticationUseCase(&memoryStorage, zerolog.Logger{})
	accountUseCase := au.NewAccountUseCase(&memoryStorage, zerolog.Logger{})

	accountTest := account.CreateAccountInput{
		Name:   "Isaac Newton",
		CPF:    "12345678910",
		Secret: "foobar",
	}
	accountCreated, _ := accountUseCase.Create(accountTest)

	t.Run("should return a token successfully", func(t *testing.T) {
		inputs := authentication.LoginInputs{
			CPF:     "12345678910",
			Secret:  "foobar",
			Account: accountCreated,
		}
		tokenGot, err := authenticationUseCase.CreateToken(inputs)
		assert.Nil(t, err)
		assert.NotEmpty(t, tokenGot)

	})

	tests := []struct {
		name      string
		inputs    authentication.LoginInputs
		wantError error
	}{
		{
			name: "should return an error because CPF is invalid",
			inputs: authentication.LoginInputs{
				CPF:     "12345678911",
				Secret:  "foobar",
				Account: accountCreated,
			},
			wantError: ErrInvalidCPF,
		},
		{
			name: "should return an error because secret is invalid",
			inputs: authentication.LoginInputs{
				CPF:     "12345678910",
				Secret:  "foo",
				Account: accountCreated,
			},
			wantError: ErrInvalidSecret,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenGot, err := authenticationUseCase.CreateToken(tt.inputs)

			assert.Equal(t, tt.wantError, err)
			assert.Empty(t, tokenGot)

		})
	}
}
