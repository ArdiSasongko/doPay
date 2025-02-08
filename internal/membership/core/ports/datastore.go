package ports

import "context"

type CacheRepositoryAdapter interface {
	GetUserSessionFromCache(context.Context)
	UpdateUserSessionIntoCache(context.Context)
}

type DatabaseRepositoryAdapter interface {
	InsertUserInfoIntoDB(context.Context)
	GetUserInfoFromDB(context.Context)
	GetUserByUsername(context.Context)
}
