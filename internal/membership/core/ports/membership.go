package ports

import "context"

type MembershipServiceAdapter interface {
	Register(context.Context)
	GetUserInfo(context.Context)
	SubmitLogin(context.Context)
	SubmitLogout(context.Context)
}
