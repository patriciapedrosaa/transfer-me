package usecase

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	fakeID      = "deef3e68-a8ed-4320-b09c-037eab340125"
	wrongFakeID = "deef3e68-a8ed-4320-b09c-037eab340125"
)

var fakeAccountGetBalanceTests = entities.Account{
	AccountID: fakeID,
	Name:      "Sirius Black",
	CPF:       "12345678911",
	Balance:   100,
}

func TestGetBalance(t *testing.T) {
	tests := []struct {
		name          string
		repositoryErr error
		id            string
		wantErr       error
		wantResult    int
	}{
		{
			name:          "should return balance successfully",
			repositoryErr: nil,
			id:            fakeID,
			wantErr:       nil,
			wantResult:    100,
		},
		{
			name:          "should return an error because id is not found",
			repositoryErr: ErrNotFound,
			id:            wrongFakeID,
			wantErr:       ErrNotFound,
			wantResult:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeGetBalanceAccountRepository(tt.repositoryErr)
			accountUseCase := NewAccountUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := accountUseCase.GetBalance(ctx, tt.id)

			assert.Equal(t, tt.wantResult, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func generateFakeGetBalanceAccountRepository(err error) account.RepositoryMock {
	if err != nil {
		return account.RepositoryMock{
			GetByIdFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, err
			},
		}
	}
	return account.RepositoryMock{
		GetByIdFunc: func(ctx context.Context, id string) (entities.Account, error) {
			return fakeAccountGetBalanceTests, nil
		},
	}
}
