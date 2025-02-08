package main

import (
	"fmt"

	"github.com/ArdiSasongko/doPay/internal/payment/adapter/handler"
	"github.com/ArdiSasongko/doPay/internal/payment/adapter/repository"
	"github.com/ArdiSasongko/doPay/internal/payment/core/services"
	"github.com/ArdiSasongko/doPay/lib/config"
	"github.com/spf13/viper"
)

func initConfig(configName string) (*config.Config, error) {
	cfg := &config.Config{}
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	fmt.Printf("%#v\n", cfg)
	return cfg, err
}

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg)
	hdl, err := handler.NewHandler(cfg, services.NewService(cfg, repo))
	if err != nil {
		return nil, err
	}

	return hdl, nil
}
