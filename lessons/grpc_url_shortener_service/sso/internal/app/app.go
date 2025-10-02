package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.NewApp(log, grpcPort, storagePath, tokenTTL)

	return &App{
		GRPCServer: grpcApp,
	}
}
