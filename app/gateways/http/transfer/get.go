package transfer

import (
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
	"time"
)

type GetTransferResponse struct {
	TransferID           string    `json:"id" validate:"required"`
	Amount               int       `json:"amount" validate:"required"`
	AccountOriginID      string    `json:"account_origin_id" validate:"required"`
	AccountDestinationID string    `json:"account_destination_id" validate:"required"`
	CreatedAt            time.Time `json:"created_at" validate:"required"`
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accountID := r.Context().Value(http_server.AccountID).(string)

	transferList, err := h.useCase.GetTransfersByAccountID(ctx, accountID)
	if err != nil {
		h.logger.Err(err).
			Int("status_code", http.StatusBadRequest).
			Msg("error occurred when try get transfers")
		http_server.ResponseError(w, http.StatusBadRequest, ErrUnexpected)
		return
	}

	response := make([]GetTransferResponse, len(transferList))
	for index, transfer := range transferList {
		response[index].TransferID = transfer.TransferID
		response[index].Amount = transfer.Amount
		response[index].AccountOriginID = transfer.AccountOriginID
		response[index].AccountDestinationID = transfer.AccountDestinationID
		response[index].CreatedAt = transfer.CreatedAt
	}
	h.logger.Info().
		Int("status_code", http.StatusOK).
		Int("total transfers listed", len(transferList)).
		Msg("transfers were successfully listed")

	http_server.ResponseSuccess(w, http.StatusOK, response)

}
