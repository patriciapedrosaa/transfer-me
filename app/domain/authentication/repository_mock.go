// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package authentication

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			CreateTokenFunc: func(ctx context.Context, token entities.Token) error {
// 				panic("mock out the CreateToken method")
// 			},
// 			GetTokenFunc: func(ctx context.Context, id string) (entities.Token, error) {
// 				panic("mock out the GetToken method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// CreateTokenFunc mocks the CreateToken method.
	CreateTokenFunc func(ctx context.Context, token entities.Token) error

	// GetTokenFunc mocks the GetToken method.
	GetTokenFunc func(ctx context.Context, id string) (entities.Token, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateToken holds details about calls to the CreateToken method.
		CreateToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Token is the token argument value.
			Token entities.Token
		}
		// GetToken holds details about calls to the GetToken method.
		GetToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
	}
	lockCreateToken sync.RWMutex
	lockGetToken    sync.RWMutex
}

// CreateToken calls CreateTokenFunc.
func (mock *RepositoryMock) CreateToken(ctx context.Context, token entities.Token) error {
	callInfo := struct {
		Ctx   context.Context
		Token entities.Token
	}{
		Ctx:   ctx,
		Token: token,
	}
	mock.lockCreateToken.Lock()
	mock.calls.CreateToken = append(mock.calls.CreateToken, callInfo)
	mock.lockCreateToken.Unlock()
	if mock.CreateTokenFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.CreateTokenFunc(ctx, token)
}

// CreateTokenCalls gets all the calls that were made to CreateToken.
// Check the length with:
//     len(mockedRepository.CreateTokenCalls())
func (mock *RepositoryMock) CreateTokenCalls() []struct {
	Ctx   context.Context
	Token entities.Token
} {
	var calls []struct {
		Ctx   context.Context
		Token entities.Token
	}
	mock.lockCreateToken.RLock()
	calls = mock.calls.CreateToken
	mock.lockCreateToken.RUnlock()
	return calls
}

// GetToken calls GetTokenFunc.
func (mock *RepositoryMock) GetToken(ctx context.Context, id string) (entities.Token, error) {
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
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
	return mock.GetTokenFunc(ctx, id)
}

// GetTokenCalls gets all the calls that were made to GetToken.
// Check the length with:
//     len(mockedRepository.GetTokenCalls())
func (mock *RepositoryMock) GetTokenCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetToken.RLock()
	calls = mock.calls.GetToken
	mock.lockGetToken.RUnlock()
	return calls
}