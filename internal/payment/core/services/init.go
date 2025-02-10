package services

import (
	"github.com/ArdiSasongko/doPay/internal/payment/core/ports"
	"github.com/ArdiSasongko/doPay/lib/config"
)

type PaymentService struct {
	config     *config.Config
	repository ports.WalletRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.WalletRepositoryAdapter) ports.PaymentServiceAdapter {
	return &PaymentService{
		config:     cfg,
		repository: repository,
	}
}
