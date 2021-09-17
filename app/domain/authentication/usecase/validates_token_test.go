package usecase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestValidatesToken(t *testing.T) {
	tests := []struct {
		name           string
		errCreateToken error
		errGetToken    error
		token          string
		wantErr        error
		wantResult     entities.Token
	}{
		{
			name:           "Should return a token successfully",
			errCreateToken: nil,
			errGetToken:    nil,
			token:          generateFakeToken(),
			wantErr:        nil,
			wantResult:     fakeToken,
		},
		{
			name:           "Should return error because token signature method is invalid",
			errCreateToken: nil,
			errGetToken:    nil,
			token:          "eyJhbGciOiJQUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.VMOk4ckLBFstkXMK0FuApNcG2FqNjs0_D8YubBKDOJ09IQV5XEexJBUv9YYkf60JBphZw_puMMEYlOGzlvgTCNeVCmzCDTPG2mvyuUG80ZPM-3B_uZyt23TbHKNF5GFvDa0X3Fa-aXTrM4cwjMVSku0YEbTKNvN1Ei3tyuQaPEWFRG-0Z6X_7ATSDYjrhmOk-RKP6dj5Yd2f4xMPf1ab4u9u98HFHBubmXR0dl9HmnVPfOwGCn0DuA9YqfG_NEzDaUVFTWsoBIajDYDZSbtHFp-D5ylE3WbomkaYjxpkZAAHAXyXwExW1QM3FM_JZZhmMywuOuIa0gZJAwOUvXuoyg",
			wantErr:        ErrMethodInvalid,
			wantResult:     entities.Token{},
		},
		{
			name:           "Should return error because token expired",
			errCreateToken: nil,
			errGetToken:    nil,
			token:          generateExpiredToken(fakeToken.Name, fakeAccount.AccountID),
			wantErr:        errors.New("Token is expired"),
			wantResult:     entities.Token{},
		},
		{
			name:           "Should return error because token is not found",
			errCreateToken: nil,
			errGetToken:    ErrTokenNotFound,
			token:          generateFakeToken(),
			wantErr:        ErrTokenNotFound,
			wantResult:     entities.Token{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeAuthenticationRepository(tt.errCreateToken, tt.errGetToken)
			authenticationUseCase := NewAuthenticationUseCase(&repository, "", zerolog.Logger{})
			ctx := context.Background()

			got, err := authenticationUseCase.ValidatesToken(ctx, tt.token)

			if tt.wantErr == nil {
				assert.NotEmpty(t, got)
				assert.NotEmpty(t, got.ID)
				assert.NotEmpty(t, got.IssuedAt)
				assert.NotEmpty(t, got.ExpiredAt)
				assert.Equal(t, fakeAccount.AccountID, got.Subject)
				assert.Equal(t, fakeAccount.Name, got.Name)
				assert.Equal(t, got.Issuer, entities.ISSUER)
				assert.Empty(t, err)
			} else {
				assert.Empty(t, got)
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
			}

		})
	}

}

func generateExpiredToken(username string, subject string) string {
	token := entities.Token{
		ID:       uuid.New().String(),
		Name:     username,
		Subject:  subject,
		Issuer:   entities.ISSUER,
		IssuedAt: time.Time{},
	}
	token.ExpiredAt = token.IssuedAt.Add(time.Minute * 20)

	atClaims := jwt.MapClaims{
		"id":   token.ID,
		"name": token.Name,
		"sub":  token.Subject,
		"iss":  token.Issuer,
		"iat":  token.IssuedAt.Unix(),
		"exp":  token.ExpiredAt.Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return accessTokenString
}

func generateFakeToken() string {
	token := entities.Token{
		ID:       fakeToken.ID,
		Name:     fakeToken.Name,
		Subject:  fakeToken.Subject,
		Issuer:   entities.ISSUER,
		IssuedAt: time.Now(),
	}
	token.ExpiredAt = token.IssuedAt.Add(time.Minute * entities.DURATION)

	atClaims := jwt.MapClaims{
		"id":   token.ID,
		"name": token.Name,
		"sub":  token.Subject,
		"iss":  token.Issuer,
		"iat":  token.IssuedAt.Unix(),
		"exp":  token.ExpiredAt.Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return accessTokenString
}
