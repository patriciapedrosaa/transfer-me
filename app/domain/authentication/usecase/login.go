package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var (
	ErrInvalidCPF    = errors.New("cpf is wrong")
	ErrInvalidSecret = errors.New("secret is wrong")
)

func (a Authentication) checkLogin(inputs authentication.LoginInputs) (bool, error) {
	a.logger.Info().
		Str("account_CPF terminated in", inputs.CPF[len(inputs.CPF)-3:]).
		Msg("starting check login CPF.")
	isValidCPF := inputs.CPF == string(inputs.Account.CPF)
	if !isValidCPF {
		a.logger.Error().
			Err(ErrInvalidCPF).
			Str("account_CPF terminated in", inputs.CPF[len(inputs.CPF)-3:]).
			Msg("error occurred when was validating the CPF")
		return false, ErrInvalidCPF
	}
	err := vos.CompareHashAndSecret(inputs.Secret, string(inputs.Account.Secret))
	if err != nil {
		a.logger.Error().Err(err).Msg("error occurred when was validating secret")
		return false, ErrInvalidSecret
	}
	a.logger.Info().
		Str("account_CPF terminated in", inputs.CPF[len(inputs.CPF)-3:]).
		Msg("login was validated with success!")
	return true, err
}
