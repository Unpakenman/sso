package main

import (
	"context"
	"sso/internal/app/config"
	appLog "sso/internal/app/log"
	bootstrap "sso/internal/bootstrap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	logger := appLog.SetupLogger(cfg.Namespace)
	bootstrap.RunService(ctx, cfg, logger)
}
