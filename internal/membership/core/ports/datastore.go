package ports

import (
	"context"

	"github.com/ArdiSasongko/doPay/internal/membership/core/models"
)

type DatastoreRepositoryAdapter interface {
	// redis
	GetUserSessionFromCache(context.Context)
	UpdateUserSessionIntoCache(context.Context)

	// postgres
	InsertUserInfoIntoDB(context.Context)
	GetUserInfoFromDB(context.Context, string) (models.UserProfileInfo, error)
	GetUserByUsername(context.Context)
}
