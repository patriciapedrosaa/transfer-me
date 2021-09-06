package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var fakeAccount = entities.Account{
	AccountID: uuid.New().String(),
	Name:      "Isac Newton",
	CPF:       "12345678910",
	Secret:    "$2a$04$mGAXaUMA5zHNKeRDXBB9/eab5toJzS/I3lR.BAMA5raQbHdaUCtXG",
	Balance:   100,
	CreatedAt: time.Now(),
}

var fakeToken = entities.Token{
	ID:        uuid.New().String(),
	Name:      fakeAccount.Name,
	Subject:   fakeAccount.AccountID,
	Issuer:    entities.ISSUER,
	IssuedAt:  time.Now().UTC(),
	ExpiredAt: time.Now().UTC().Add(time.Minute * 15),
}

var unexpectedErr = errors.New("unexpected error to create token")

func TestCreateToken(t *testing.T) {
	tests := []struct {
		name           string
		errCreateToken error
		errGetToken    error
		inputs         authentication.LoginInputs
		wantError      error
	}{
		{
			name:           "should return a token successfully",
			errCreateToken: nil,
			errGetToken:    nil,
			inputs: authentication.LoginInputs{
				CPF:     "12345678910",
				Secret:  "mySecret",
				Account: fakeAccount,
			},
			wantError: nil,
		},
		{
			name:           "should return an error because CPF is invalid",
			errCreateToken: nil,
			errGetToken:    nil,
			inputs: authentication.LoginInputs{
				CPF:     "12345678911",
				Secret:  "mySecret",
				Account: fakeAccount,
			},
			wantError: ErrInvalidCPF,
		},
		{
			name:           "should return an error because secret is invalid",
			errCreateToken: nil,
			errGetToken:    nil,
			inputs: authentication.LoginInputs{
				CPF:     "12345678910",
				Secret:  "foo",
				Account: fakeAccount,
			},
			wantError: ErrInvalidSecret,
		},
		{
			name:           "should return an error because could not create a token",
			errCreateToken: unexpectedErr,
			errGetToken:    nil,
			inputs: authentication.LoginInputs{
				CPF:     "12345678910",
				Secret:  "mySecret",
				Account: fakeAccount,
			},
			wantError: unexpectedErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeAuthenticationRepository(tt.errCreateToken, tt.errGetToken)
			authenticationUseCase := NewAuthenticationUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := authenticationUseCase.CreateToken(ctx, tt.inputs)

			if tt.wantError == nil {
				assert.Nil(t, err)
				assert.NotEmpty(t, got)
			} else {
				assert.Equal(t, tt.wantError, err)
				assert.Empty(t, got)
			}

		})
	}
}

func generateFakeAuthenticationRepository(errCreateToken error, errGetToken error) authentication.RepositoryMock {
	if errCreateToken != nil {
		return authentication.RepositoryMock{
			CreateTokenFunc: func(ctx context.Context, token entities.Token) error {
				return errCreateToken
			},
		}
	}
	if errGetToken != nil {
		return authentication.RepositoryMock{
			GetTokenFunc: func(ctx context.Context, id string) (entities.Token, error) {
				return entities.Token{}, errGetToken
			},
		}
	}
	return authentication.RepositoryMock{
		CreateTokenFunc: func(ctx context.Context, token entities.Token) error {
			return nil
		},
		GetTokenFunc: func(ctx context.Context, id string) (entities.Token, error) {
			return fakeToken, nil
		},
	}
}
