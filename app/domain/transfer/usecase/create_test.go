package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	fakeOriginAccount = entities.Account{
		AccountID: uuid.New().String(),
		Name:      "Julius",
		CPF:       "12345678910",
		Secret:    "secret",
		Balance:   100,
		CreatedAt: time.Now(),
	}

	fakeDestinationAccount = entities.Account{
		AccountID: uuid.New().String(),
		Name:      "Rochelle Rock",
		CPF:       "12345678911",
		Secret:    "secret",
		Balance:   100,
		CreatedAt: time.Now(),
	}
)

func TestCreateTransfer(t *testing.T) {
	tests := []struct {
		name              string
		errCreateTransfer error
		inputs            transfer.CreateTransferInput
		wantErr           error
		wantResult        entities.Transfer
	}{
		{
			name:              "should return a transfer successfully",
			errCreateTransfer: nil,
			inputs: transfer.CreateTransferInput{
				OriginAccount:      fakeOriginAccount,
				DestinationAccount: fakeDestinationAccount,
				Amount:             50,
			},
			wantErr: nil,
			wantResult: entities.Transfer{
				AccountOriginID:      fakeOriginAccount.AccountID,
				AccountDestinationID: fakeDestinationAccount.AccountID,
				Amount:               50,
			},
		},
		{
			name:              "should return an error because amount is zero",
			errCreateTransfer: entities.ErrInvalidAmount,
			inputs: transfer.CreateTransferInput{
				OriginAccount:      fakeOriginAccount,
				DestinationAccount: fakeDestinationAccount,
				Amount:             0,
			},
			wantErr:    entities.ErrInvalidAmount,
			wantResult: entities.Transfer{},
		},
		{
			name:              "should return an error because amount is negative",
			errCreateTransfer: entities.ErrInvalidAmount,
			inputs: transfer.CreateTransferInput{
				OriginAccount:      fakeOriginAccount,
				DestinationAccount: fakeDestinationAccount,
				Amount:             -10,
			},
			wantErr:    entities.ErrInvalidAmount,
			wantResult: entities.Transfer{},
		},
		{
			name:              "should return an error because accounts are equals",
			errCreateTransfer: entities.ErrInvalidDestinationAccount,
			inputs: transfer.CreateTransferInput{
				OriginAccount:      fakeOriginAccount,
				DestinationAccount: fakeOriginAccount,
				Amount:             50,
			},
			wantErr:    entities.ErrInvalidDestinationAccount,
			wantResult: entities.Transfer{},
		},
		{
			name:              "should return an error because insufficient funds",
			errCreateTransfer: entities.ErrInvalidTransfer,
			inputs: transfer.CreateTransferInput{
				OriginAccount:      fakeOriginAccount,
				DestinationAccount: fakeDestinationAccount,
				Amount:             500,
			},
			wantErr:    entities.ErrInvalidTransfer,
			wantResult: entities.Transfer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeTransferRepository(tt.errCreateTransfer, nil)
			transferUseCase := NewTransferUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := transferUseCase.Create(ctx, tt.inputs)

			assert.Equal(t, tt.wantResult.AccountOriginID, got.AccountOriginID)
			assert.Equal(t, tt.wantResult.AccountDestinationID, got.AccountDestinationID)
			assert.Equal(t, tt.wantResult.Amount, got.Amount)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func generateFakeTransferRepository(errCreateTransfer, errGetTransfer error) transfer.RepositoryMock {
	if errCreateTransfer != nil {
		return transfer.RepositoryMock{
			CreateTransferFunc: func(ctx context.Context, transfer entities.Transfer, accountID string) error {
				return errCreateTransfer
			},
		}
	}
	if errGetTransfer != nil {
		return transfer.RepositoryMock{
			GetTransfersByAccountIDFunc: func(ctx context.Context, accountID string) ([]entities.Transfer, error) {
				return []entities.Transfer{}, errGetTransfer
			},
		}
	}
	return transfer.RepositoryMock{
		CreateTransferFunc: func(ctx context.Context, transfer entities.Transfer, accountID string) error {
			return nil
		},
		GetTransfersByAccountIDFunc: func(ctx context.Context, accountID string) ([]entities.Transfer, error) {
			return []entities.Transfer{}, nil
		},
	}
}
