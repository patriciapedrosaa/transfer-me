package usecase

import (
	"context"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

var (
	ErrUnexpected = errors.New("something went wrong")
)

func (t Transfer) Create(ctx context.Context, input transfer.CreateTransferInput) (entities.Transfer, error) {
	log := t.logger.With().
		Str("account_ID", input.OriginAccount.AccountID).
		Str("account_ID", input.DestinationAccount.AccountID).
		Int("amount", input.Amount).
		Logger()

	log.Info().Msg("creating transfer.")
	newTransfer, err := entities.NewTransfer(input.OriginAccount, input.DestinationAccount, input.Amount)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to create transfer")
		return entities.Transfer{}, err
	}
	err = t.transferRepository.CreateTransfer(ctx, newTransfer, input.OriginAccount.AccountID)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to create transfer")
		return entities.Transfer{}, ErrUnexpected
	}
	log.Info().
		Str("transfer_ID", newTransfer.TransferID).
		Msg("transfer created with success!")
	return newTransfer, nil
}
