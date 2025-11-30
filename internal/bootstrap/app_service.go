package bootstrap

import (
	"context"
	"log/slog"
	"sso/internal/app/client/pg"
	"sso/internal/app/config"
	"time"
)

func RunService(ctx context.Context, cfg *config.Values, logger *slog.Logger) {

	db, err := pg.NewClient(ctx, 5, *cfg.DbPayments)
	if err != nil {
		logger.Error("failed to connect to DB", slog.String("error", err.Error()))
		return
	}

	if err := db.Ping(ctx); err != nil {
		logger.Error("db ping failed", slog.String("error", err.Error()))
		return
	}

	app, err := New(logger, cfg.GRPCServer.Port, cfg.TokenLifeTime)
	if err != nil {
		logger.Error("app init failed", slog.String("error", err.Error()))
		return
	}

	// Graceful shutdown
	go func() {
		if err := app.GRPCServer.Run(); err != nil {
			logger.Error("grpc server error", slog.String("error", err.Error()))
		}
	}()

	<-ctx.Done()

	logger.Info("shutting down...")
	app.GRPCServer.Stop()
	db.Close()
}

type App struct {
	GRPCServer *GrpcApp
}

func New(
	log *slog.Logger,
	grpcPort int32,
	tokenLT time.Duration,
) (*App, error) {

	grpcApp := NewGrpcServer(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
