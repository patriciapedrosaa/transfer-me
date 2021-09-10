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

var (
	fakeAccount1 = entities.Account{
		AccountID: uuid.New().String(),
		Name:      "Ronald Weasley",
		CPF:       "12345678910",
		Secret:    "foobar",
		Balance:   100,
		CreatedAt: time.Now(),
	}
	fakeAccount2 = entities.Account{
		AccountID: uuid.New().String(),
		Name:      "Ginevra Weasley",
		CPF:       "12345678911",
		Secret:    "foobar",
		Balance:   100,
		CreatedAt: time.Now(),
	}
)

func TestGetAccounts(t *testing.T) {
	tests := []struct {
		name          string
		repositoryErr error
		wantErr       error
		wantResult    []entities.Account
	}{
		{
			name:          "should return a list of account",
			repositoryErr: nil,
			wantErr:       nil,
			wantResult: []entities.Account{
				{
					Name:    "Ronald Weasley",
					CPF:     "12345678910",
					Balance: 100,
				},
				{
					Name:    "Ginevra Weasley",
					CPF:     "12345678911",
					Balance: 100,
				},
			},
		},
		{
			name:          "should return an error because repository returned an unexpected err",
			repositoryErr: unexpectedRepositoryErr,
			wantErr:       nil,
			wantResult:    []entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeGetAccountRepository(tt.repositoryErr)
			accountUseCase := NewAccountUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := accountUseCase.GetAccounts(ctx)

			for k := range got {
				assert.Equal(t, tt.wantResult[k].Name, got[k].Name)
				assert.Equal(t, tt.wantResult[k].Name, got[k].Name)
				assert.Equal(t, tt.wantResult[k].CPF, got[k].CPF)
				assert.Equal(t, tt.wantResult[k].Balance, got[k].Balance)
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func generateFakeGetAccountRepository(err error) account.RepositoryMock {
	if err != nil {
		return account.RepositoryMock{
			GetByCpfFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, err
			},
		}
	}
	return account.RepositoryMock{
		GetAccountsFunc: func(ctx context.Context) ([]entities.Account, error) {
			return []entities.Account{
				fakeAccount1,
				fakeAccount2,
			}, nil
		},
	}
}
