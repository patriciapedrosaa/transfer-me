package usecase

import (
	"errors"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransfer(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage)
	accountUsecase := au.NewAccountUsecase(&memoryStorage)

	createAccountInput1 := au.CreateAccountInput{
		Name:   "John Locke",
		CPF:    "12345678910",
		Secret: "foobar",
	}

	createAccountInput2 := au.CreateAccountInput{
		Name:   "Karl Marx",
		CPF:    "12345678911",
		Secret: "foobar",
	}
	account1, _ := accountUsecase.Create(createAccountInput1)
	account2, _ := accountUsecase.Create(createAccountInput2)

	transferUsecase := NewTransferUsecase(&memoryStorage, &memoryStorage)

	fakeTransfer := CreateTransferInput{
		OriginAccountCPF:      "12345678911",
		DestinationAccountCPF: "12345678910",
		Amount:                20,
	}
	_, _ = transferUsecase.Create(fakeTransfer)
	_ = accountUsecase.UpdateBalance(vos.CPF(fakeTransfer.OriginAccountCPF), vos.CPF(fakeTransfer.DestinationAccountCPF), fakeTransfer.Amount)

	tests := []struct {
		name       string
		inputs     CreateTransferInput
		wantErr    error
		wantResult entities.Transfer
	}{
		{
			name: "should return a transfer successfully",
			inputs: CreateTransferInput{
				OriginAccountCPF:      string(account1.CPF),
				DestinationAccountCPF: string(account2.CPF),
				Amount:                50,
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
			inputs: CreateTransferInput{
				OriginAccountCPF:      "12345678911",
				DestinationAccountCPF: "12345678910",
				Amount:                10,
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
			inputs: CreateTransferInput{
				OriginAccountCPF:      "12345678915",
				DestinationAccountCPF: string(account2.CPF),
				Amount:                50,
			},
			wantErr:    errors.New("not found"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because amount is zero",
			inputs: CreateTransferInput{
				OriginAccountCPF:      string(account1.CPF),
				DestinationAccountCPF: string(account2.CPF),
				Amount:                0,
			},
			wantErr:    errors.New("the amount must be greater than zero"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because amount is negative",
			inputs: CreateTransferInput{
				OriginAccountCPF:      string(account1.CPF),
				DestinationAccountCPF: string(account2.CPF),
				Amount:                -10,
			},
			wantErr:    errors.New("the amount must be greater than zero"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because accounts are equals",
			inputs: CreateTransferInput{
				OriginAccountCPF:      string(account1.CPF),
				DestinationAccountCPF: string(account1.CPF),
				Amount:                50,
			},
			wantErr:    errors.New("accounts must be different"),
			wantResult: entities.Transfer{},
		},
		{
			name: "should return an error because insufficient funds",
			inputs: CreateTransferInput{
				OriginAccountCPF:      string(account1.CPF),
				DestinationAccountCPF: string(account2.CPF),
				Amount:                500,
			},
			wantErr:    errors.New("insufficient funds"),
			wantResult: entities.Transfer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transferCreated, err := transferUsecase.Create(tt.inputs)

			assert.Equal(t, tt.wantResult.AccountOriginID, transferCreated.AccountOriginID)
			assert.Equal(t, tt.wantResult.AccountDestinationID, transferCreated.AccountDestinationID)
			assert.Equal(t, tt.wantResult.Amount, transferCreated.Amount)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
