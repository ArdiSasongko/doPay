package main

import (
	"os"
	"path/filepath"

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

	configPath := filepath.Join("..", "..", "files", "config")
	cfg, err := config.InitConfig("config.membership", configPath)
	if err != nil {
		log.Fatalf("failed to parse config, %#v", err)
	}

	startService(cfg)
}
