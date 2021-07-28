package usecase

import (
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateToken(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, nil, authenticationStorage)
	authenticationUsecase := NewAuthenticationUseCase(&memoryStorage, &memoryStorage)
	accountUsecase := au.NewAccountUsecase(&memoryStorage)

	accountTest := au.CreateAccountInput{
		Name:   "Isaac Newton",
		CPF:    "12345678910",
		Secret: "foobar",
	}
	accountCreated, _ := accountUsecase.Create(accountTest)

	t.Run("should return a token successfully", func(t *testing.T) {
		inputs := LoginInputs{
			CPF:     "12345678910",
			Secret:  "foobar",
			Account: accountCreated,
		}
		tokenGot, err := authenticationUsecase.CreateToken(inputs)
		assert.Nil(t, err)
		assert.NotEmpty(t, tokenGot)

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
			tokenGot, err := authenticationUsecase.CreateToken(tt.inputs)

			assert.Equal(t, tt.wantError, err)
			assert.Empty(t, tokenGot)

		})
	}
}
