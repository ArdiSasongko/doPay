package repository

import (
	"context"
	"fmt"

	"github.com/ArdiSasongko/doPay/internal/membership/core/models"
)

func (d *DatastoreRepository) GetUserSessionFromCache(context.Context) {
}

func (d *DatastoreRepository) UpdateUserSessionIntoCache(context.Context) {
}

func (d *DatastoreRepository) InsertUserInfoIntoDB(context.Context) {
}

func (d *DatastoreRepository) GetUserInfoFromDB(ctx context.Context, accountNumber string) (data models.UserProfileInfo, err error) {
	err = d.queries.GetUserInfo.GetContext(ctx, &data, map[string]interface{}{
		"account_number": accountNumber,
	})
	if err != nil {
		fmt.Printf("err : %#v\n", err)
	}
	return
}

func (d *DatastoreRepository) GetUserByUsername(context.Context) {
}
