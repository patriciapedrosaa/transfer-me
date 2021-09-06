package usecase

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateBalance(t *testing.T) {
	tests := []struct {
		name               string
		errGetAccount      error
		errUpdateBalance   error
		originAccount      entities.Account
		destinationAccount entities.Account
		amount             int
		wantErr            error
	}{
		{
			name:               "should update balances successfully",
			errUpdateBalance:   nil,
			errGetAccount:      nil,
			amount:             10,
			originAccount:      fakeAccount1,
			destinationAccount: fakeAccount2,
			wantErr:            nil,
		},
		{
			name:               "should return an error because account was not found",
			errGetAccount:      ErrNotFound,
			errUpdateBalance:   nil,
			originAccount:      fakeAccount1,
			destinationAccount: fakeAccount2,
			amount:             100,
			wantErr:            ErrNotFound,
		},
		{
			name:               "should return an error because could not update balance",
			errGetAccount:      nil,
			errUpdateBalance:   unexpectedRepositoryErr,
			originAccount:      fakeAccount1,
			destinationAccount: fakeAccount2,
			amount:             100,
			wantErr:            unexpectedRepositoryErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeUpdateBalanceRepository(tt.errGetAccount, tt.errUpdateBalance)
			accountUseCase := NewAccountUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			err := accountUseCase.UpdateBalance(ctx, tt.originAccount.AccountID, tt.destinationAccount.AccountID, tt.amount)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func generateFakeUpdateBalanceRepository(repositoryGetErr error, repositoryUpdateErr error) account.RepositoryMock {
	if repositoryGetErr != nil {
		return account.RepositoryMock{
			CreateAccountFunc: nil,
			GetByIdFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return fakeAccount, repositoryGetErr
			},
		}
	}
	if repositoryUpdateErr != nil {
		return account.RepositoryMock{
			UpdateBalanceFunc: func(ctx context.Context, id string, value int) error {
				return repositoryUpdateErr
			},
			GetByIdFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, nil
			},
		}
	}
	return account.RepositoryMock{
		UpdateBalanceFunc: func(ctx context.Context, id string, value int) error {
			return nil
		},
	}
}
