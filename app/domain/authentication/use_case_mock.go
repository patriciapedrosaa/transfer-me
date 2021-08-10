// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package authentication

import (
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"sync"
)

// Ensure, that UseCaseMock does implement UseCase.
// If this is not the case, regenerate this file with moq.
var _ UseCase = &UseCaseMock{}

// UseCaseMock is a mock implementation of UseCase.
//
// 	func TestSomethingThatUsesUseCase(t *testing.T) {
//
// 		// make and configure a mocked UseCase
// 		mockedUseCase := &UseCaseMock{
// 			CheckLoginFunc: func(inputs LoginInputs) (bool, error) {
// 				panic("mock out the CheckLogin method")
// 			},
// 			CreateTokenFunc: func(login LoginInputs) (string, error) {
// 				panic("mock out the CreateToken method")
// 			},
// 			GetTokenFunc: func(id string) (entities.Token, error) {
// 				panic("mock out the GetToken method")
// 			},
// 			ValidatesTokenFunc: func(tokenString string) (*jwt.Token, error) {
// 				panic("mock out the ValidatesToken method")
// 			},
// 		}
//
// 		// use mockedUseCase in code that requires UseCase
// 		// and then make assertions.
//
// 	}
type UseCaseMock struct {
	// CheckLoginFunc mocks the CheckLogin method.
	CheckLoginFunc func(inputs LoginInputs) (bool, error)

	// CreateTokenFunc mocks the CreateToken method.
	CreateTokenFunc func(login LoginInputs) (string, error)

	// GetTokenFunc mocks the GetToken method.
	GetTokenFunc func(id string) (entities.Token, error)

	// ValidatesTokenFunc mocks the ValidatesToken method.
	ValidatesTokenFunc func(tokenString string) (*jwt.Token, error)

	// calls tracks calls to the methods.
	calls struct {
		// CheckLogin holds details about calls to the CheckLogin method.
		CheckLogin []struct {
			// Inputs is the inputs argument value.
			Inputs LoginInputs
		}
		// CreateToken holds details about calls to the CreateToken method.
		CreateToken []struct {
			// Login is the login argument value.
			Login LoginInputs
		}
		// GetToken holds details about calls to the GetToken method.
		GetToken []struct {
			// ID is the id argument value.
			ID string
		}
		// ValidatesToken holds details about calls to the ValidatesToken method.
		ValidatesToken []struct {
			// TokenString is the tokenString argument value.
			TokenString string
		}
	}
	lockCheckLogin     sync.RWMutex
	lockCreateToken    sync.RWMutex
	lockGetToken       sync.RWMutex
	lockValidatesToken sync.RWMutex
}

// CheckLogin calls CheckLoginFunc.
func (mock *UseCaseMock) CheckLogin(inputs LoginInputs) (bool, error) {
	callInfo := struct {
		Inputs LoginInputs
	}{
		Inputs: inputs,
	}
	mock.lockCheckLogin.Lock()
	mock.calls.CheckLogin = append(mock.calls.CheckLogin, callInfo)
	mock.lockCheckLogin.Unlock()
	if mock.CheckLoginFunc == nil {
		var (
			bOut   bool
			errOut error
		)
		return bOut, errOut
	}
	return mock.CheckLoginFunc(inputs)
}

// CheckLoginCalls gets all the calls that were made to CheckLogin.
// Check the length with:
//     len(mockedUseCase.CheckLoginCalls())
func (mock *UseCaseMock) CheckLoginCalls() []struct {
	Inputs LoginInputs
} {
	var calls []struct {
		Inputs LoginInputs
	}
	mock.lockCheckLogin.RLock()
	calls = mock.calls.CheckLogin
	mock.lockCheckLogin.RUnlock()
	return calls
}

// CreateToken calls CreateTokenFunc.
func (mock *UseCaseMock) CreateToken(login LoginInputs) (string, error) {
	callInfo := struct {
		Login LoginInputs
	}{
		Login: login,
	}
	mock.lockCreateToken.Lock()
	mock.calls.CreateToken = append(mock.calls.CreateToken, callInfo)
	mock.lockCreateToken.Unlock()
	if mock.CreateTokenFunc == nil {
		var (
			sOut   string
			errOut error
		)
		return sOut, errOut
	}
	return mock.CreateTokenFunc(login)
}

// CreateTokenCalls gets all the calls that were made to CreateToken.
// Check the length with:
//     len(mockedUseCase.CreateTokenCalls())
func (mock *UseCaseMock) CreateTokenCalls() []struct {
	Login LoginInputs
} {
	var calls []struct {
		Login LoginInputs
	}
	mock.lockCreateToken.RLock()
	calls = mock.calls.CreateToken
	mock.lockCreateToken.RUnlock()
	return calls
}

// GetToken calls GetTokenFunc.
func (mock *UseCaseMock) GetToken(id string) (entities.Token, error) {
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetToken.Lock()
	mock.calls.GetToken = append(mock.calls.GetToken, callInfo)
	mock.lockGetToken.Unlock()
	if mock.GetTokenFunc == nil {
		var (
			tokenOut entities.Token
			errOut   error
		)
		return tokenOut, errOut
	}
	return mock.GetTokenFunc(id)
}

// GetTokenCalls gets all the calls that were made to GetToken.
// Check the length with:
//     len(mockedUseCase.GetTokenCalls())
func (mock *UseCaseMock) GetTokenCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetToken.RLock()
	calls = mock.calls.GetToken
	mock.lockGetToken.RUnlock()
	return calls
}

// ValidatesToken calls ValidatesTokenFunc.
func (mock *UseCaseMock) ValidatesToken(tokenString string) (*jwt.Token, error) {
	callInfo := struct {
		TokenString string
	}{
		TokenString: tokenString,
	}
	mock.lockValidatesToken.Lock()
	mock.calls.ValidatesToken = append(mock.calls.ValidatesToken, callInfo)
	mock.lockValidatesToken.Unlock()
	if mock.ValidatesTokenFunc == nil {
		var (
			tokenOut *jwt.Token
			errOut   error
		)
		return tokenOut, errOut
	}
	return mock.ValidatesTokenFunc(tokenString)
}

// ValidatesTokenCalls gets all the calls that were made to ValidatesToken.
// Check the length with:
//     len(mockedUseCase.ValidatesTokenCalls())
func (mock *UseCaseMock) ValidatesTokenCalls() []struct {
	TokenString string
} {
	var calls []struct {
		TokenString string
	}
	mock.lockValidatesToken.RLock()
	calls = mock.calls.ValidatesToken
	mock.lockValidatesToken.RUnlock()
	return calls
}