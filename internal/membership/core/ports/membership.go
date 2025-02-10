package ports

import (
	"context"

	"github.com/ArdiSasongko/doPay/internal/membership/core/models"
)

type MembershipServiceAdapter interface {
	Register(context.Context)
	GetUserInfo(context.Context, string) (models.UserProfileInfo, error)
	SubmitLogin(context.Context, models.LoginInfo) (models.LoginResponse, error)
	SubmitLogout(context.Context)
}
