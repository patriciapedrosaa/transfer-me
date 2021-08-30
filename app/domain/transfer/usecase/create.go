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
	t.logger.Info().
		Str("account_ID", input.OriginAccount.AccountID).
		Str("account_ID", input.DestinationAccount.AccountID).
		Int("amount", input.Amount).
		Msg("creating transfer.")
	newTransfer, err := entities.NewTransfer(input.OriginAccount, input.DestinationAccount, input.Amount)
	if err != nil {
		t.logger.Error().Err(err).
			Str("account_ID", input.OriginAccount.AccountID).
			Str("account_ID", input.DestinationAccount.AccountID).
			Int("amount", input.Amount).
			Msgf("error occurred when was trying to create transfer")
		return entities.Transfer{}, err
	}
	err = t.transferRepository.CreateTransfer(newTransfer, input.OriginAccount.AccountID)
	if err != nil {
		t.logger.Error().Err(err).
			Str("account_ID", input.OriginAccount.AccountID).
			Str("account_ID", input.DestinationAccount.AccountID).
			Int("amount", input.Amount).
			Msgf("error occurred when was trying to create transfer")
		return entities.Transfer{}, ErrUnexpected
	}
	t.logger.Info().
		Str("transfer_ID", newTransfer.TransferID).
		Str("account_ID", input.OriginAccount.AccountID).
		Str("account_ID", input.DestinationAccount.AccountID).
		Int("amount", input.Amount).
		Msgf("transfer created with success!")
	return newTransfer, nil
}
