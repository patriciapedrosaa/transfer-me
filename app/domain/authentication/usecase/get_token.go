package usecase

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (a Authentication) GetToken(id string) (entities.Token, error) {
	token, err := a.authenticationRepository.GetToken(id)
	if err != nil {
		return entities.Token{}, err
	}
	return token, nil
}
