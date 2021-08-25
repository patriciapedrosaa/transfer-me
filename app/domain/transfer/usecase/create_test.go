package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransfer(t *testing.T) {
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

	fakeTransfer := transfer.CreateTransferInput{
		OriginAccountId:      account1.AccountID,
		DestinationAccountId: account2.AccountID,
		Amount:               20,
	}
	_, _ = transferUseCase.Create(fakeTransfer)
	_ = accountUseCase.UpdateBalance(account1.AccountID, account2.AccountID, fakeTransfer.Amount)

	tests := []struct {
		name       string
		inputs     transfer.CreateTransferInput
		wantErr    error
		wantResult entities.Transfer
	}{
		{
			name: "should return a transfer successfully",
			inputs: transfer.CreateTransferInput{
				OriginAccountId:      account1.AccountID,
				DestinationAccountId: account2.AccountID,
				Amount:               50,
			},
			wantErr: nil,
			wantResult: entities.Transfer{
				AccountOriginID:      account1.AccountID,
				AccountDestinationID: account2.AccountID,
				Amount:               50,
			},
		},
		{
			name: "should return another transfer successfully",
			inputs: transfer.CreateTransferInput{
				OriginAccountId:      account2.AccountID,
				DestinationAccountId: account1.AccountID,
				Amount:               10,
			},
			wantErr: nil,
			wantResult: entities.Transfer{
				AccountOriginID:      account2.AccountID,
				AccountDestinationID: account1.AccountID,
				Amount:               10,
			},
		},
		{
			name: "should return an error because account was not found",
			inputs: transfer.CreateTransferInput{
				OriginAccountId:      uuid.New().String(),
				DestinationAccountId: account2.AccountID,
				Amount:               50,
			},
			wantErr:    errors.New("not found"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because amount is zero",
			inputs: transfer.CreateTransferInput{
				OriginAccountId:      account1.AccountID,
				DestinationAccountId: account2.AccountID,
				Amount:               0,
			},
			wantErr:    errors.New("the amount must be greater than zero"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because amount is negative",
			inputs: transfer.CreateTransferInput{
				OriginAccountId:      account1.AccountID,
				DestinationAccountId: account2.AccountID,
				Amount:               -10,
			},
			wantErr:    errors.New("the amount must be greater than zero"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because accounts are equals",
			inputs: transfer.CreateTransferInput{
				OriginAccountId:      account1.AccountID,
				DestinationAccountId: account1.AccountID,
				Amount:               50,
			},
			wantErr:    errors.New("accounts must be different"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because insufficient funds",
			inputs: transfer.CreateTransferInput{
				OriginAccountId:      account1.AccountID,
				DestinationAccountId: account2.AccountID,
				Amount:               500,
			},
			wantErr:    errors.New("insufficient funds"),
			wantResult: entities.Transfer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transferCreated, err := transferUseCase.Create(tt.inputs)

			assert.Equal(t, tt.wantResult.AccountOriginID, transferCreated.AccountOriginID)
			assert.Equal(t, tt.wantResult.AccountDestinationID, transferCreated.AccountDestinationID)
			assert.Equal(t, tt.wantResult.Amount, transferCreated.Amount)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
