package services

import (
	"github.com/ArdiSasongko/doPay/internal/membership/core/ports"
	"github.com/ArdiSasongko/doPay/lib/config"
)

type MembershipService struct {
	config     *config.Config
	repository ports.DatastoreRepositoryAdapter
}

func NewMembershipService(cfg *config.Config, repository ports.DatastoreRepositoryAdapter) ports.MembershipServiceAdapter {
	return &MembershipService{
		config:     cfg,
		repository: repository,
	}
}
