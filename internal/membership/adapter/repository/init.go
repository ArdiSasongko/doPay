package repository

import (
	"github.com/ArdiSasongko/doPay/lib/config"
	"github.com/ArdiSasongko/doPay/lib/datastore/postgres"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type DatastoreRepository struct {
	config   *config.Config
	client   *redis.Client
	dbClient postgres.DBInterface
	queries  statementQueries
}

func NewDatastoreRepository(cfg *config.Config, client *redis.Client, dbClient postgres.DBInterface) *DatastoreRepository {
	repo := &DatastoreRepository{
		config:   cfg,
		client:   client,
		dbClient: dbClient,
	}
	if err := repo.prepareStatements(); err != nil {
		log.Error(err)
	}

	return repo
}
