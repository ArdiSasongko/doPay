package handler

import (
	"context"

	"github.com/ArdiSasongko/doPay/internal/membership/core/models"
	"github.com/ArdiSasongko/doPay/lib/protos/v1/membership"
)

func (h Handler) GetUserInfo(ctx context.Context, req *membership.UserInfoPayload) (*membership.UserInfoResponse, error) {
	dataUser, err := h.membership.GetUserInfo(ctx, req.AccountNumber)
	if err != nil {
		return nil, err
	}
	return &membership.UserInfoResponse{
		Email:         dataUser.Email,
		Fullname:      dataUser.Fullname,
		AccountNumber: dataUser.AccountNumber,
		Status:        dataUser.Status,
	}, nil
}

func (h Handler) Register(ctx context.Context, req *membership.RegisterRequest) (*membership.RegisterResponse, error) {
	return nil, nil
}

func (h Handler) SubmitLogin(ctx context.Context, req *membership.LoginRequest) (*membership.LoginResponse, error) {
	h.membership.SubmitLogin(ctx, models.LoginInfo{
		Username: req.Username,
		Password: req.Password,
	})
	return nil, nil
}

func (h Handler) SubmitLogout(ctx context.Context, req *membership.LogoutRequest) (*membership.LogoutResponse, error) {
	return nil, nil
}
