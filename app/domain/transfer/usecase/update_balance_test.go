package usecase

import (
	"errors"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateBalance(t *testing.T) {
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string]memory.Transfer)
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
	fakeAccount := entities.Account{
		AccountID: "11223344556",
		Name:      "fake",
		CPF:       "11223344556",
		Secret:    "secret",
		Balance:   100,
	}

	transferUsecase := NewTransferUsecase(&memoryStorage, &memoryStorage)

	transfer := CreateTransferInput{
		OriginAccountCPF:      string(account1.CPF),
		DestinationAccountCPF: string(account2.CPF),
		Amount:                50,
	}
	_, _ = transferUsecase.Create(transfer)

	tests := []struct {
		name               string
		originAccount      entities.Account
		destinationAccount entities.Account
		amount             int
		wantErr            error
	}{
		{
			name:               "should update balances successfully",
			amount:             100,
			originAccount:      account1,
			destinationAccount: account2,
			wantErr:            nil,
		},
		{
			name:               "should return an error because account was not found",
			amount:             50,
			originAccount:      fakeAccount,
			destinationAccount: account1,
			wantErr:            errors.New("not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := transferUsecase.UpdateBalance(tt.originAccount, tt.destinationAccount, tt.amount)
			accountOrigin, _ := accountUsecase.GetByCpf(string(account1.CPF))
			accountDestiny, _ := accountUsecase.GetByCpf(string(account2.CPF))

			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				assert.Equal(t, accountOrigin.Balance, 0)
				assert.Equal(t, accountDestiny.Balance, 200)
			}
		})
	}

}
