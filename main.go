package main

import (
	au "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	auth "github.com/patriciapedrosaa/transfer-me/app/domain/authentication/usecase"
	tu "github.com/patriciapedrosaa/transfer-me/app/domain/transfer/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/db/memory"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/http/account"
	auth_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http/authentication"
	transfer_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http/transfer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	apiPort = ":8000"
)

func main() {
	//Log
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	//Memory
	accountStorage := make(map[string]memory.Account)
	transferStorage := make(map[string][]memory.Transfer)
	authenticationStorage := make(map[string]memory.Token)
	memoryStorage := memory.NewMemoryStorage(accountStorage, transferStorage, authenticationStorage)

	//Use Cases
	accountUseCase := au.NewAccountUseCase(&memoryStorage, logger)
	transferUseCase := tu.NewTransferUseCase(&memoryStorage, logger)
	authenticationUseCase := auth.NewAuthenticationUseCase(&memoryStorage, logger)

	//Routes
	accountHandler := account.NewHandler(accountUseCase, logger)
	authHandler := auth_server.NewHandler(authenticationUseCase, accountUseCase, logger)
	transferHandler := transfer_server.NewHandler(transferUseCase, accountUseCase, logger)
	api := http.NewApi(accountHandler, authHandler, transferHandler, logger)
	log.Info().Msgf("Starting api at port %s", apiPort)
	api.Start(apiPort)
}
