package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var fakeAccountGetByID = entities.Account{
	AccountID: fakeID,
	Name:      "Dino Thomas",
	CPF:       "12345678911",
	Secret:    "foobar",
	Balance:   100,
	CreatedAt: time.Time{},
}

func TestGetById(t *testing.T) {
	tests := []struct {
		name          string
		repositoryErr error
		id            string
		wantErr       error
		wantResult    entities.Account
	}{
		{
			name:          "should return an account successfully",
			repositoryErr: nil,
			id:            fakeID,
			wantErr:       nil,
			wantResult:    fakeAccountGetByID,
		},
		{
			name:          "should return err invalid id",
			repositoryErr: nil,
			id:            invalidFakeID,
			wantErr:       ErrInvalidId,
			wantResult:    entities.Account{},
		},
		{
			name:          "should return err not found",
			repositoryErr: ErrNotFound,
			id:            uuid.New().String(),
			wantErr:       ErrNotFound,
			wantResult:    entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeGetByIDRepository(tt.repositoryErr)
			accountUseCase := NewAccountUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := accountUseCase.GetById(ctx, tt.id)

			assert.Equal(t, tt.wantResult.Name, got.Name)
			assert.Equal(t, tt.wantResult.CPF, got.CPF)
			assert.Equal(t, tt.wantResult.Balance, got.Balance)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func generateFakeGetByIDRepository(err error) account.RepositoryMock {
	if err != nil {
		return account.RepositoryMock{
			GetByIdFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, err
			},
		}
	}
	return account.RepositoryMock{
		GetByIdFunc: func(ctx context.Context, id string) (entities.Account, error) {
			return fakeAccountGetByID, nil
		},
	}
}
