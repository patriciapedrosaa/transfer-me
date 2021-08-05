package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTransfers(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage, nil)
	accountUsecase := au.NewAccountUseCase(&memoryStorage)

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
	account1, _ := accountUsecase.Create(createAccountInput1)
	account2, _ := accountUsecase.Create(createAccountInput2)

	transferUsecase := NewTransferUsecase(&memoryStorage, &memoryStorage)

	transfer1 := CreateTransferInput{
		OriginAccountId:      account1.AccountID,
		DestinationAccountId: account2.AccountID,
		Amount:               50,
	}
	transfer2 := CreateTransferInput{
		OriginAccountId:      account1.AccountID,
		DestinationAccountId: account2.AccountID,
		Amount:               10,
	}

	_, _ = transferUsecase.Create(transfer1)
	_, _ = transferUsecase.Create(transfer2)

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
			transfers, err := transferUsecase.GetTransfersByAccountID(account1.AccountID)

			for k, _ := range transfers {
				assert.Equal(t, tt.wantResult[k].AccountOriginID, transfers[k].AccountOriginID)
				assert.Equal(t, tt.wantResult[k].AccountDestinationID, transfers[k].AccountDestinationID)
				assert.Equal(t, tt.wantResult[k].Amount, transfers[k].Amount)
			}
			assert.Equal(t, tt.wantErr, err)

		})
	}
}
