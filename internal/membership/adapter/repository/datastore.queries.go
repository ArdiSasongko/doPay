package repository

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmoiron/sqlx"
)

const (
	GetUserInfo = `
		select account_number, fullname, email, status
		from user_profile p
		where account_number = :account_number
	`
)

type statementQueries struct {
	GetUserInfo       *sqlx.NamedStmt
	GetUserByUsername *sqlx.NamedStmt
	InsertUserAuth    *sqlx.NamedStmt
	RegisterUserInfo  *sqlx.NamedStmt
}

func (r *DatastoreRepository) prepareStatements() error {
	// get user info
	stmtNamed, err := r.dbClient.PrepareNamed(GetUserInfo)
	if err != nil {
		log.Error("prepare statement GetUserInfo err, ", err.Error())
		return err
	}
	r.queries.GetUserInfo = stmtNamed

	return nil
}
