package entities

import (
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
	"time"
)

const (
	DURATION = 15
	ISSUER   = "JWT"
)

type Login struct {
	CPF    vos.CPF
	Secret vos.Secret
}

type Token struct {
	ID        string
	Name      string
	Subject   string
	Issuer    string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewCreateToken(username, subject string) (Token, error) {
	token := Token{
		ID:       uuid.New().String(),
		Name:     username,
		Subject:  subject,
		Issuer:   ISSUER,
		IssuedAt: time.Time{},
	}
	token.ExpiredAt = token.IssuedAt.Add(time.Minute * DURATION)

	return token, nil
}
