package main

import (
	"log"

	"github.com/ArdiSasongko/doPay/internal/membership/adapter/handler"
	"github.com/ArdiSasongko/doPay/internal/membership/adapter/repository"
	"github.com/ArdiSasongko/doPay/internal/membership/core/services"
	"github.com/ArdiSasongko/doPay/lib/config"
	"github.com/ArdiSasongko/doPay/lib/datastore/postgres"
	"github.com/go-redis/redis/v8"
)

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	redis := InitRedis(cfg)
	dbClient, _ := InitDB(cfg)

	repo := repository.NewDatastoreRepository(cfg, redis, dbClient)
	hdl, err := handler.NewHandler(cfg, services.NewMembershipService(cfg, repo))
	if err != nil {
		return nil, err
	}

	return hdl, nil
}

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}

func InitDB(cfg *config.Config) (postgres.DBInterface, error) {
	dbConn := postgres.InitDB(cfg)
	client, err := dbConn.InitiateConnect()
	if err != nil {
		log.Fatal("failed to connect db, err :%w", err)
	}

	return client, nil
}
