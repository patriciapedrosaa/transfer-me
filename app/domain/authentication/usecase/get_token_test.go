package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetToken(t *testing.T) {
	tests := []struct {
		name           string
		errCreateToken error
		errGetToken    error
		tokenID        string
		wantError      error
	}{
		{
			name:           "Should return a token successfully",
			errCreateToken: nil,
			errGetToken:    nil,
			tokenID:        fakeToken.ID,
			wantError:      nil,
		},
		{
			name:           "Should return err invalid id",
			errCreateToken: nil,
			errGetToken:    ErrInvalidId,
			tokenID:        "12dae-cec4-4c18-aa-4e4a1bb59",
			wantError:      ErrInvalidId,
		},
		{
			name:           "Should return err not found",
			errCreateToken: nil,
			errGetToken:    ErrTokenNotFound,
			tokenID:        uuid.New().String(),
			wantError:      ErrTokenNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := generateFakeAuthenticationRepository(tt.errCreateToken, tt.errGetToken)
			authenticationUseCase := NewAuthenticationUseCase(&repository, zerolog.Logger{})
			ctx := context.Background()

			got, err := authenticationUseCase.getToken(ctx, tt.tokenID)

			if tt.wantError == nil {
				assert.Equal(t, got.Subject, fakeToken.Subject)
				assert.Equal(t, got.Issuer, fakeToken.Issuer)
				assert.Equal(t, got.Name, fakeToken.Name)
				assert.Empty(t, err)
			} else {
				assert.Error(t, err)
				assert.Empty(t, got)
				assert.Equal(t, err, tt.wantError)

			}
		})
	}
}
