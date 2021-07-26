package authentication

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

type Repository interface {
	CreateToken(token entities.Token) error
	GetToken(id string) (entities.Token, error)
}
