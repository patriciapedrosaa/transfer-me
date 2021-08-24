package transfer

import (
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
	"time"
)

type GetTransferResponse struct {
	TransferID string    `json:"id" validate:"required"`
	CreatedAt  time.Time `json:"created_at" validate:"required"`
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	accountID := r.Context().Value(http_server.ContextID).(string)

	transferList, err := h.useCase.GetTransfersByAccountID(accountID)
	if err != nil {
		http_server.ResponseError(w, http.StatusForbidden, ErrUnexpected)
		return
	}

	response := make([]GetTransferResponse, len(transferList))
	for index, transfer := range transferList {
		response[index].TransferID = transfer.TransferID
		response[index].CreatedAt = transfer.CreatedAt
	}

	http_server.ResponseSuccess(w, http.StatusOK, response)

}
