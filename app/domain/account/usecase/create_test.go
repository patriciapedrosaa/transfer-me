package usecase

import (
	"context"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	fakeAccount = entities.Account{
		Name: "Percy Jackson",
		CPF:  "12345678913",
	}

	unexpectedRepositoryErr = errors.New("unexpected repository error")
)

func TestAccountCreate(t *testing.T) {
	tests := []struct {
		name                string
		repositoryCreateErr error
		repositoryGetErr    error
		inputs              account.CreateAccountInput
		wantErr             error
		wantResult          entities.Account
	}{
		{
			name:                "should creates an account successfully",
			repositoryCreateErr: nil,
			repositoryGetErr:    nil,
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
			name:                "should return an error because name is invalid",
			repositoryCreateErr: nil,
			repositoryGetErr:    nil,
			inputs: account.CreateAccountInput{
				Name:   "",
				CPF:    "12345678911",
				Secret: "foobar",
			},
			wantErr:    errors.New("invalid name"),
			wantResult: entities.Account{},
		},
		{
			name:                "should return an error because cpf is invalid",
			repositoryCreateErr: nil,
			repositoryGetErr:    nil,
			inputs: account.CreateAccountInput{
				Name:   "Grover Underwood",
				CPF:    "123456789",
				Secret: "foobar",
			},
			wantErr:    errors.New("invalid cpf"),
			wantResult: entities.Account{},
		},
		{
			name:                "should return an error because secret is invalid",
			repositoryCreateErr: nil,
			repositoryGetErr:    nil,
			inputs: account.CreateAccountInput{
				Name:   "Grover Underwood",
				CPF:    "12345678912",
				Secret: "",
			},
			wantErr:    errors.New("invalid secret"),
			wantResult: entities.Account{},
		},
		{
			name:                "should return an error because account already exist",
			repositoryCreateErr: nil,
			repositoryGetErr:    ErrAlreadyExist,
			inputs: account.CreateAccountInput{
				Name:   "Percy Jackson",
				CPF:    "12345678913",
				Secret: "foobar",
			},
			wantErr:    errors.New("account already exist"),
			wantResult: entities.Account{},
		},
		{
			name:                "should return an error because repository returned an unexpected err",
			repositoryCreateErr: unexpectedRepositoryErr,
			repositoryGetErr:    nil,
			inputs: account.CreateAccountInput{
				Name:   "Percy Jackson",
				CPF:    "12345678913",
				Secret: "foobar",
			},
			wantErr:    errors.New("unexpected repository error"),
			wantResult: entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeAccountRepository(tt.repositoryGetErr, tt.repositoryCreateErr)
			accountUseCase := NewAccountUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := accountUseCase.Create(ctx, tt.inputs)

			assert.Equal(t, tt.wantResult.Name, got.Name)
			assert.Equal(t, tt.wantResult.CPF, got.CPF)
			assert.Equal(t, tt.wantResult.Balance, got.Balance)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
			}

			if tt.wantErr == nil {
				assert.NotEqual(t, tt.inputs.Secret, got.Secret)
			}

		})
	}

}

func generateFakeAccountRepository(repositoryGetErr error, repositoryCreateErr error) account.RepositoryMock {
	if repositoryGetErr != nil {
		return account.RepositoryMock{
			CreateAccountFunc: nil,
			GetByCpfFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return fakeAccount, repositoryGetErr
			},
		}
	}
	if repositoryCreateErr != nil {
		return account.RepositoryMock{
			CreateAccountFunc: func(ctx context.Context, account entities.Account) error {
				return repositoryCreateErr
			},
			GetByCpfFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, nil
			},
		}
	}
	return account.RepositoryMock{
		CreateAccountFunc: func(ctx context.Context, account entities.Account) error {
			return nil
		},
	}
}
