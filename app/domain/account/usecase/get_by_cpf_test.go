package usecase

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fakeAccountGetByCPF = entities.Account{
	Name:    "Dino Thomas",
	CPF:     "12345678911",
	Secret:  "foobar",
	Balance: 100,
}

func TestGetByCpf(t *testing.T) {
	tests := []struct {
		name          string
		repositoryErr error
		cpf           string
		wantErr       error
		wantResult    entities.Account
	}{
		{
			name:          "should return an account successfully",
			repositoryErr: nil,
			cpf:           "12345678911",
			wantErr:       nil,
			wantResult:    fakeAccountGetByCPF,
		},
		{
			name:          "should return err not found",
			repositoryErr: ErrNotFound,
			cpf:           "12345678922",
			wantErr:       ErrNotFound,
			wantResult:    entities.Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeGetByCPFRepository(tt.repositoryErr)
			accountUseCase := NewAccountUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := accountUseCase.GetByCpf(ctx, tt.cpf)

			assert.Equal(t, tt.wantResult.Name, got.Name)
			assert.Equal(t, tt.wantResult.CPF, got.CPF)
			assert.Equal(t, tt.wantResult.Balance, got.Balance)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func generateFakeGetByCPFRepository(err error) account.RepositoryMock {
	if err != nil {
		return account.RepositoryMock{
			GetByCpfFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, err
			},
		}
	}
	return account.RepositoryMock{
		GetByCpfFunc: func(ctx context.Context, id string) (entities.Account, error) {
			return fakeAccountGetByCPF, nil
		},
	}
}
