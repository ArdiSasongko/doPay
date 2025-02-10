package ports

import (
	"context"

	"github.com/ArdiSasongko/doPay/internal/payment/core/models"
)

// PaymentServiceAdapter - port primary
type PaymentServiceAdapter interface {
	TransferUserBalance(ctx context.Context, payload models.TransferBalancePayload) (float64, error)
}
