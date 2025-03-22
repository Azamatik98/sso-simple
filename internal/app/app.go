package app

import (
	"github.com/Azamatik98/sso/internal/app/grpc"
	"github.com/Azamatik98/sso/internal/services/auth"
	"github.com/Azamatik98/sso/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	gRPCPort int,
	storagePath string,
	tokenTTl time.Duration,
) *App {

	storage, err := sqlite.NewStorage(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.NewAuth(log, storage, storage, storage, tokenTTl)

	grpcApp := grpcapp.NewApp(log, authService, gRPCPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
