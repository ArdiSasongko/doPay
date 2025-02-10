package handler

import (
	"github.com/ArdiSasongko/doPay/internal/payment/core/ports"
	"github.com/ArdiSasongko/doPay/lib/config"
	pbHandler "github.com/ArdiSasongko/doPay/lib/protos/v1/payment"
)

type Handler struct {
	pbHandler.UnimplementedPaymentServiceServer
	config         *config.Config
	paymentService ports.PaymentServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.PaymentServiceAdapter) (*Handler, error) {
	return &Handler{
		config:         cfg,
		paymentService: adapter,
	}, nil
}
