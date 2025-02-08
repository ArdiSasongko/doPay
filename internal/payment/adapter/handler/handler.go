package handler

import (
	"context"

	"github.com/ArdiSasongko/doPay/internal/payment/core/models"
	"github.com/ArdiSasongko/doPay/lib/protos/v1/payment"
)

func (h *Handler) TransferBalanceService(ctx context.Context, in *payment.TransferBalanceRequest) (*payment.TransferBalanceResponse, error) {
	amount, err := h.paymentService.TransferUserBalance(ctx, models.TransferBalancePayload{
		SourceUserID: in.GetSourceUserId(),
		TargetUserID: in.GetDestination(),
		Amount:       in.GetAmount(),
	})

	return &payment.TransferBalanceResponse{
		Success:           err == nil,
		DestinationAmount: amount,
	}, nil
}
