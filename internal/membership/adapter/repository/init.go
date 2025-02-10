package repository

import (
	"github.com/ArdiSasongko/doPay/lib/config"
	"github.com/ArdiSasongko/doPay/lib/datastore/postgres"
	"github.com/go-redis/redis/v8"
)

type DatastoreRepository struct {
	config   *config.Config
	client   *redis.Client
	dbClient postgres.DBInterface
}

func NewDatastoreRepository(cfg *config.Config, client *redis.Client, dbClient postgres.DBInterface) *DatastoreRepository {
	return &DatastoreRepository{
		config:   cfg,
		client:   client,
		dbClient: dbClient,
	}
}
