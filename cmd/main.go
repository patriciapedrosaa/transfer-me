package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/patriciapedrosaa/transfer-me/app/common/configuration"
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	auth "github.com/patriciapedrosaa/transfer-me/app/domain/authentication/usecase"
	tu "github.com/patriciapedrosaa/transfer-me/app/domain/transfer/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/postgres"
	account_repository "github.com/patriciapedrosaa/transfer-me/app/gateways/db/postgres/account"
	authentication_repository "github.com/patriciapedrosaa/transfer-me/app/gateways/db/postgres/authentication"
	transfer_repository "github.com/patriciapedrosaa/transfer-me/app/gateways/db/postgres/transfer"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/http/account"
	auth_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http/authentication"
	transfer_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http/transfer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := configuration.LoadConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to load app configuration")
	}

	loglevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to load log level")
	}
	zerolog.SetGlobalLevel(loglevel)

	pgConn := connectDB(logger, cfg.Postgres)
	accountRepository := account_repository.NewRepository(pgConn)
	transferRepository := transfer_repository.NewRepository(pgConn)
	authenticationRepository := authentication_repository.NewRepository(pgConn)
	defer pgConn.Close()

	err = postgres.RunMigrations(cfg.Postgres.URL())
	if err != nil {
		logger.Fatal().Msg("error running postgres migrations")
	}

	accountUseCase := au.NewAccountUseCase(accountRepository, logger)
	transferUseCase := tu.NewTransferUseCase(transferRepository, logger)
	authenticationUseCase := auth.NewAuthenticationUseCase(authenticationRepository, cfg.TokenKey, logger)

	accountHandler := account.NewHandler(accountUseCase, logger)
	authHandler := auth_server.NewHandler(authenticationUseCase, accountUseCase, logger)
	transferHandler := transfer_server.NewHandler(transferUseCase, accountUseCase, logger)
	api := http.NewApi(accountHandler, authHandler, transferHandler, logger)
	log.Info().Msgf("Starting api at port %s", cfg.API.Port)
	api.Start(cfg.API.Port)
}

func connectDB(logger zerolog.Logger, postgres configuration.PostgresConfig) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), postgres.DSN())
	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to connect to database")
	}
	return pool
}
