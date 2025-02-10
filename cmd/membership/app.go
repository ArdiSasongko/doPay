package main

import (
	"os"

	"github.com/ArdiSasongko/doPay/lib/config"
	"github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	"github.com/go-kratos/kratos/v2/log"
	zlog "github.com/rs/zerolog"
)

func main() {
	zlogT := zlog.New(os.Stdout).With().CallerWithSkipFrameCount(4).Timestamp().Logger()
	logger := zerolog.NewLogger(&zlogT)
	log.SetLogger(logger)

	log.Info("Starting backend service")

	cfg, err := config.InitConfig("config.membership")
	if err != nil {
		log.Fatalf("failed to parse config, %#v", err)
	}

	startService(cfg)
}
