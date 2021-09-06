package usecase

import (
	"context"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTransfers(t *testing.T) {
	tests := []struct {
		name           string
		errGetTransfer error
		wantErr        error
		wantResult     []entities.Transfer
	}{
		{
			name:           "should return a list of transfers successfully",
			errGetTransfer: nil,
			wantErr:        nil,
			wantResult: []entities.Transfer{
				{
					AccountOriginID:      fakeOriginAccount.AccountID,
					AccountDestinationID: fakeDestinationAccount.AccountID,
					Amount:               50,
				},
				{
					AccountOriginID:      fakeOriginAccount.AccountID,
					AccountDestinationID: fakeDestinationAccount.AccountID,
					Amount:               10,
				},
			},
		},
		{
			name:           "should return err not found",
			errGetTransfer: errors.New("not found"),
			wantErr:        errors.New("not found"),
			wantResult:     []entities.Transfer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeTransferRepository(nil, tt.errGetTransfer)
			transferUseCase := NewTransferUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := transferUseCase.GetTransfersByAccountID(ctx, fakeOriginAccount.AccountID)

			for k, _ := range got {
				assert.Equal(t, tt.wantResult[k].AccountOriginID, got[k].AccountOriginID)
				assert.Equal(t, tt.wantResult[k].AccountDestinationID, got[k].AccountDestinationID)
				assert.Equal(t, tt.wantResult[k].Amount, got[k].Amount)
			}
			assert.Equal(t, tt.wantErr, err)

		})
	}
}
