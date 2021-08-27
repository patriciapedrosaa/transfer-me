package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

var (
	ErrUnexpected = errors.New("something went wrong")
)

func (t Transfer) Create(input transfer.CreateTransferInput) (entities.Transfer, error) {
	t.logger.Info().Msgf("Creating account from origin account: %s to destination account: %s in the amount of: %d...", input.OriginAccount.AccountID, input.DestinationAccount.AccountID, input.Amount)
	newTransfer, err := entities.NewTransfer(input.OriginAccount, input.DestinationAccount, input.Amount)
	if err != nil {
		t.logger.Error().Err(err).Msgf("Occurred when was trying to create transfer from origin account: %s to destination account: %s in the amount of: %d.", input.OriginAccount.AccountID, input.DestinationAccount.AccountID, input.Amount)
		return entities.Transfer{}, err
	}
	err = t.transferRepository.CreateTransfer(newTransfer, input.OriginAccount.AccountID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("Occurred when was trying to create transfer from origin account: %s to destination account: %s in the amount of: %d.", input.OriginAccount.AccountID, input.DestinationAccount.AccountID, input.Amount)
		return entities.Transfer{}, ErrUnexpected
	}
	t.logger.Info().Msgf("Transfer created with success from origin account: %s to destination account: %s in the amount of: %d. Transfer ID: %s", input.OriginAccount.AccountID, input.DestinationAccount.AccountID, input.Amount, newTransfer.TransferID)
	return newTransfer, nil
}
