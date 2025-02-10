package handler

import (
	"github.com/ArdiSasongko/doPay/internal/membership/core/ports"
	"github.com/ArdiSasongko/doPay/lib/config"
	pbHandler "github.com/ArdiSasongko/doPay/lib/protos/v1/membership"
)

type Handler struct {
	pbHandler.UnimplementedMembershipServiceServer
	config     *config.Config
	membership ports.MembershipServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.MembershipServiceAdapter) (*Handler, error) {
	return &Handler{
		config:     cfg,
		membership: adapter,
	}, nil
}
