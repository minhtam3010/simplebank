package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/minhtam3010/simplebank/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"fromAccountId" binding:"required,min=1"`
	ToAccountID   int64  `json:"toAccountId" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR"`
}

type TransferHandler struct {
	store db.Store
}

func NewTransferHandler(store db.Store) *TransferHandler {
	return &TransferHandler{store: store}
}

func (th *TransferHandler) CreateTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTransfersParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        int64(req.Amount),
	}

	account, err := th.store.CreateTransfers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
