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
	a.logger.Info().Msgf("Checking login for CPF: %s...", inputs.CPF)
	isValidCPF := inputs.CPF == string(inputs.Account.CPF)
	if !isValidCPF {
		a.logger.Error().Err(ErrInvalidCPF).Msgf("Occurred when was validating the CPF: %s", inputs.CPF)
		return false, ErrInvalidCPF
	}
	err := vos.CompareHashAndSecret(inputs.Secret, string(inputs.Account.Secret))
	if err != nil {
		a.logger.Error().Err(err).Msgf("Occurred when was validating secret for CPF: %s", inputs.CPF)
		return false, ErrInvalidSecret
	}
	a.logger.Info().Msgf("Login was validated with success for CPF: %s", inputs.CPF)
	return true, err
}
