package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidAmount             = errors.New("the amount must be greater than zero")
	ErrInvalidTransfer           = errors.New("insufficient funds")
	ErrInvalidDestinationAccount = errors.New("accounts must be different")
)

type Transfer struct {
	TransferID           string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(originAccount, destinationAccount Account, amount int) (Transfer, error) {
	if amount <= 0 {
		return Transfer{}, ErrInvalidAmount
	}

	if equalsAccounts(originAccount, destinationAccount) {
		return Transfer{}, ErrInvalidDestinationAccount
	}

	if !isValidTransfer(originAccount.Balance, amount) {
		return Transfer{}, ErrInvalidTransfer
	}

	return Transfer{
		TransferID:           uuid.New().String(),
		AccountOriginID:      originAccount.AccountID,
		AccountDestinationID: destinationAccount.AccountID,
		Amount:               amount,
		CreatedAt:            time.Now().UTC(),
	}, nil
}

func isValidTransfer(balance int, amount int) bool {
	if balance < amount {
		return false
	}
	return true
}

func equalsAccounts(accountOrigin, accountDestination Account) bool {
	return accountOrigin.CPF == accountDestination.CPF
}
