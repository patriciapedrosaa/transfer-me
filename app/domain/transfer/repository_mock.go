// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package transfer

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
// 			CreateTransferFunc: func(ctx context.Context, transfer entities.Transfer, accountID string) error {
// 				panic("mock out the CreateTransfer method")
// 			},
// 			GetTransfersByAccountIDFunc: func(ctx context.Context, accountID string) ([]entities.Transfer, error) {
// 				panic("mock out the GetTransfersByAccountID method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// CreateTransferFunc mocks the CreateTransfer method.
	CreateTransferFunc func(ctx context.Context, transfer entities.Transfer, accountID string) error

	// GetTransfersByAccountIDFunc mocks the GetTransfersByAccountID method.
	GetTransfersByAccountIDFunc func(ctx context.Context, accountID string) ([]entities.Transfer, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateTransfer holds details about calls to the CreateTransfer method.
		CreateTransfer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transfer is the transfer argument value.
			Transfer entities.Transfer
			// AccountID is the accountID argument value.
			AccountID string
		}
		// GetTransfersByAccountID holds details about calls to the GetTransfersByAccountID method.
		GetTransfersByAccountID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AccountID is the accountID argument value.
			AccountID string
		}
	}
	lockCreateTransfer          sync.RWMutex
	lockGetTransfersByAccountID sync.RWMutex
}

// CreateTransfer calls CreateTransferFunc.
func (mock *RepositoryMock) CreateTransfer(ctx context.Context, transfer entities.Transfer, accountID string) error {
	callInfo := struct {
		Ctx       context.Context
		Transfer  entities.Transfer
		AccountID string
	}{
		Ctx:       ctx,
		Transfer:  transfer,
		AccountID: accountID,
	}
	mock.lockCreateTransfer.Lock()
	mock.calls.CreateTransfer = append(mock.calls.CreateTransfer, callInfo)
	mock.lockCreateTransfer.Unlock()
	if mock.CreateTransferFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.CreateTransferFunc(ctx, transfer, accountID)
}

// CreateTransferCalls gets all the calls that were made to CreateTransfer.
// Check the length with:
//     len(mockedRepository.CreateTransferCalls())
func (mock *RepositoryMock) CreateTransferCalls() []struct {
	Ctx       context.Context
	Transfer  entities.Transfer
	AccountID string
} {
	var calls []struct {
		Ctx       context.Context
		Transfer  entities.Transfer
		AccountID string
	}
	mock.lockCreateTransfer.RLock()
	calls = mock.calls.CreateTransfer
	mock.lockCreateTransfer.RUnlock()
	return calls
}

// GetTransfersByAccountID calls GetTransfersByAccountIDFunc.
func (mock *RepositoryMock) GetTransfersByAccountID(ctx context.Context, accountID string) ([]entities.Transfer, error) {
	callInfo := struct {
		Ctx       context.Context
		AccountID string
	}{
		Ctx:       ctx,
		AccountID: accountID,
	}
	mock.lockGetTransfersByAccountID.Lock()
	mock.calls.GetTransfersByAccountID = append(mock.calls.GetTransfersByAccountID, callInfo)
	mock.lockGetTransfersByAccountID.Unlock()
	if mock.GetTransfersByAccountIDFunc == nil {
		var (
			transfersOut []entities.Transfer
			errOut       error
		)
		return transfersOut, errOut
	}
	return mock.GetTransfersByAccountIDFunc(ctx, accountID)
}

// GetTransfersByAccountIDCalls gets all the calls that were made to GetTransfersByAccountID.
// Check the length with:
//     len(mockedRepository.GetTransfersByAccountIDCalls())
func (mock *RepositoryMock) GetTransfersByAccountIDCalls() []struct {
	Ctx       context.Context
	AccountID string
} {
	var calls []struct {
		Ctx       context.Context
		AccountID string
	}
	mock.lockGetTransfersByAccountID.RLock()
	calls = mock.calls.GetTransfersByAccountID
	mock.lockGetTransfersByAccountID.RUnlock()
	return calls
}