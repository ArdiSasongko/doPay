package ports

import "context"

type DatastoreRepositoryAdapter interface {
	// redis
	GetUserSessionFromCache(context.Context)
	UpdateUserSessionIntoCache(context.Context)

	// postgres
	InsertUserInfoIntoDB(context.Context)
	GetUserInfoFromDB(context.Context)
	GetUserByUsername(context.Context)
}
