package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTransfers(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage, nil)
	accountUseCase := au.NewAccountUseCase(&memoryStorage)

	createAccountInput1 := account.CreateAccountInput{
		Name:   "John Locke",
		CPF:    "12345678910",
		Secret: "foobar",
	}

	createAccountInput2 := account.CreateAccountInput{
		Name:   "Karl Marx",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	account1, _ := accountUseCase.Create(createAccountInput1)
	account2, _ := accountUseCase.Create(createAccountInput2)

	transferUseCase := NewTransferUseCase(&memoryStorage)

	transfer1 := transfer.CreateTransferInput{
		OriginAccountId:      account1.AccountID,
		DestinationAccountId: account2.AccountID,
		Amount:               50,
	}
	transfer2 := transfer.CreateTransferInput{
		OriginAccountId:      account1.AccountID,
		DestinationAccountId: account2.AccountID,
		Amount:               10,
	}

	_, _ = transferUseCase.Create(transfer1)
	_, _ = transferUseCase.Create(transfer2)

	tests := []struct {
		name       string
		wantErr    error
		wantResult []entities.Transfer
	}{
		{
			name:    "should return a list of transfers successfully",
			wantErr: nil,
			wantResult: []entities.Transfer{
				{
					AccountOriginID:      account1.AccountID,
					AccountDestinationID: account2.AccountID,
					Amount:               50,
				},
				{
					AccountOriginID:      account1.AccountID,
					AccountDestinationID: account2.AccountID,
					Amount:               10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transfers, err := transferUseCase.GetTransfersByAccountID(account1.AccountID)

			for k, _ := range transfers {
				assert.Equal(t, tt.wantResult[k].AccountOriginID, transfers[k].AccountOriginID)
				assert.Equal(t, tt.wantResult[k].AccountDestinationID, transfers[k].AccountDestinationID)
				assert.Equal(t, tt.wantResult[k].Amount, transfers[k].Amount)
			}
			assert.Equal(t, tt.wantErr, err)

		})
	}
}
