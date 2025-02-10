package services

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/ArdiSasongko/doPay/internal/membership/core/models"
)

func (s *MembershipService) Register(context.Context) {}

func (s *MembershipService) GetUserInfo(ctx context.Context, accountNumber string) (models.UserProfileInfo, error) {
	dbUser, err := s.repository.GetUserInfoFromDB(ctx, accountNumber)
	if err != nil {
		log.Err(err).Msg("GetUserInfo.GetUserInfoFromDB")
		return dbUser, err
	}
	return dbUser, nil
}

func (s *MembershipService) SubmitLogin(ctx context.Context, request models.LoginInfo) (models.LoginResponse, error) {
	return models.LoginResponse{}, nil
}

func (s *MembershipService) SubmitLogout(context.Context) {}
