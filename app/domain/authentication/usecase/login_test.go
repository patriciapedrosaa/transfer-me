package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckLogin(t *testing.T) {
	repository := generateFakeAuthenticationRepository(nil, nil)
	authenticationUseCase := NewAuthenticationUseCase(&repository, zerolog.Logger{})
	tests := []struct {
		name      string
		inputs    authentication.LoginInputs
		wantError error
	}{
		{
			name: "should return valid login successfully",
			inputs: authentication.LoginInputs{
				CPF:     "12345678910",
				Secret:  "mySecret",
				Account: fakeAccount,
			},
			wantError: nil,
		},
		{
			name: "should return an error because CPF is invalid",
			inputs: authentication.LoginInputs{
				CPF:     "12345678911",
				Secret:  "foobar",
				Account: fakeAccount,
			},
			wantError: ErrInvalidCPF,
		},
		{
			name: "should return an error because secret is invalid",
			inputs: authentication.LoginInputs{
				CPF:     "12345678910",
				Secret:  "foo",
				Account: fakeAccount,
			},
			wantError: ErrInvalidSecret,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authenticationUseCase.checkLogin(tt.inputs)

			if tt.wantError == nil {
				assert.Nil(t, err)
				assert.True(t, got)
			} else {
				assert.Equal(t, tt.wantError, err)
				assert.False(t, got)
			}
		})
	}

}
