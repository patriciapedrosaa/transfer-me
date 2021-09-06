package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
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

const (
	apiPort    = ":8000"
	host       = "localhost"
	user       = "postgres"
	dbPassword = "postgres"
	dbName     = "transfer-me"
	dbPort     = 5432
)

func main() {
	//Log
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	//Memory

	//accountStorage := make(map[string]memory.Account)
	//transferStorage := make(map[string][]memory.Transfer)
	//authenticationStorage := make(map[string]memory.Token)
	//memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage, authenticationStorage)

	//Database
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, dbPassword, host, dbPort, dbName)
	pgConn := connectDB(databaseUrl)
	accountRepository := account_repository.NewRepository(pgConn)
	transferRepository := transfer_repository.NewRepository(pgConn)
	authenticationRepository := authentication_repository.NewRepository(pgConn)
	defer pgConn.Close()

	//Migrations
	err := postgres.RunMigrations(databaseUrl)
	if err != nil {
		log.Fatal().Msg("error running postgres migrations")
	}

	//Use Cases
	accountUseCase := au.NewAccountUseCase(accountRepository, logger)
	transferUseCase := tu.NewTransferUseCase(transferRepository, logger)
	authenticationUseCase := auth.NewAuthenticationUseCase(authenticationRepository, logger)

	//Routes
	accountHandler := account.NewHandler(accountUseCase, logger)
	authHandler := auth_server.NewHandler(authenticationUseCase, accountUseCase, logger)
	transferHandler := transfer_server.NewHandler(transferUseCase, accountUseCase, logger)
	api := http.NewApi(accountHandler, authHandler, transferHandler, logger)
	log.Info().Msgf("Starting api at port %s", apiPort)
	api.Start(apiPort)
}

func connectDB(databaseUrl string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}
	return pool
}
