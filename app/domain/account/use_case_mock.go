// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package account

import (
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
// 			CreateFunc: func(input CreateAccountInput) (entities.Account, error) {
// 				panic("mock out the Create method")
// 			},
// 			GetAccountsFunc: func() ([]entities.Account, error) {
// 				panic("mock out the GetAccounts method")
// 			},
// 			GetBalanceFunc: func(id string) (int, error) {
// 				panic("mock out the GetBalance method")
// 			},
// 			GetByIdFunc: func(id string) (entities.Account, error) {
// 				panic("mock out the GetById method")
// 			},
// 			UpdateBalanceFunc: func(originAccountId string, destinationAccountId string, amount int) error {
// 				panic("mock out the UpdateBalance method")
// 			},
// 		}
//
// 		// use mockedUseCase in code that requires UseCase
// 		// and then make assertions.
//
// 	}
type UseCaseMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(input CreateAccountInput) (entities.Account, error)

	// GetAccountsFunc mocks the GetAccounts method.
	GetAccountsFunc func() ([]entities.Account, error)

	// GetBalanceFunc mocks the GetBalance method.
	GetBalanceFunc func(id string) (int, error)

	// GetByIdFunc mocks the GetById method.
	GetByIdFunc func(id string) (entities.Account, error)

	// UpdateBalanceFunc mocks the UpdateBalance method.
	UpdateBalanceFunc func(originAccountId string, destinationAccountId string, amount int) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Input is the input argument value.
			Input CreateAccountInput
		}
		// GetAccounts holds details about calls to the GetAccounts method.
		GetAccounts []struct {
		}
		// GetBalance holds details about calls to the GetBalance method.
		GetBalance []struct {
			// ID is the id argument value.
			ID string
		}
		// GetById holds details about calls to the GetById method.
		GetById []struct {
			// ID is the id argument value.
			ID string
		}
		// UpdateBalance holds details about calls to the UpdateBalance method.
		UpdateBalance []struct {
			// OriginAccountId is the originAccountId argument value.
			OriginAccountId string
			// DestinationAccountId is the destinationAccountId argument value.
			DestinationAccountId string
			// Amount is the amount argument value.
			Amount int
		}
	}
	lockCreate        sync.RWMutex
	lockGetAccounts   sync.RWMutex
	lockGetBalance    sync.RWMutex
	lockGetById       sync.RWMutex
	lockUpdateBalance sync.RWMutex
}

// Create calls CreateFunc.
func (mock *UseCaseMock) Create(input CreateAccountInput) (entities.Account, error) {
	callInfo := struct {
		Input CreateAccountInput
	}{
		Input: input,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	if mock.CreateFunc == nil {
		var (
			accountOut entities.Account
			errOut     error
		)
		return accountOut, errOut
	}
	return mock.CreateFunc(input)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedUseCase.CreateCalls())
func (mock *UseCaseMock) CreateCalls() []struct {
	Input CreateAccountInput
} {
	var calls []struct {
		Input CreateAccountInput
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// GetAccounts calls GetAccountsFunc.
func (mock *UseCaseMock) GetAccounts() ([]entities.Account, error) {
	callInfo := struct {
	}{}
	mock.lockGetAccounts.Lock()
	mock.calls.GetAccounts = append(mock.calls.GetAccounts, callInfo)
	mock.lockGetAccounts.Unlock()
	if mock.GetAccountsFunc == nil {
		var (
			accountsOut []entities.Account
			errOut      error
		)
		return accountsOut, errOut
	}
	return mock.GetAccountsFunc()
}

// GetAccountsCalls gets all the calls that were made to GetAccounts.
// Check the length with:
//     len(mockedUseCase.GetAccountsCalls())
func (mock *UseCaseMock) GetAccountsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAccounts.RLock()
	calls = mock.calls.GetAccounts
	mock.lockGetAccounts.RUnlock()
	return calls
}

// GetBalance calls GetBalanceFunc.
func (mock *UseCaseMock) GetBalance(id string) (int, error) {
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetBalance.Lock()
	mock.calls.GetBalance = append(mock.calls.GetBalance, callInfo)
	mock.lockGetBalance.Unlock()
	if mock.GetBalanceFunc == nil {
		var (
			nOut   int
			errOut error
		)
		return nOut, errOut
	}
	return mock.GetBalanceFunc(id)
}

// GetBalanceCalls gets all the calls that were made to GetBalance.
// Check the length with:
//     len(mockedUseCase.GetBalanceCalls())
func (mock *UseCaseMock) GetBalanceCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetBalance.RLock()
	calls = mock.calls.GetBalance
	mock.lockGetBalance.RUnlock()
	return calls
}

// GetById calls GetByIdFunc.
func (mock *UseCaseMock) GetById(id string) (entities.Account, error) {
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetById.Lock()
	mock.calls.GetById = append(mock.calls.GetById, callInfo)
	mock.lockGetById.Unlock()
	if mock.GetByIdFunc == nil {
		var (
			accountOut entities.Account
			errOut     error
		)
		return accountOut, errOut
	}
	return mock.GetByIdFunc(id)
}

// GetByIdCalls gets all the calls that were made to GetById.
// Check the length with:
//     len(mockedUseCase.GetByIdCalls())
func (mock *UseCaseMock) GetByIdCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetById.RLock()
	calls = mock.calls.GetById
	mock.lockGetById.RUnlock()
	return calls
}

// UpdateBalance calls UpdateBalanceFunc.
func (mock *UseCaseMock) UpdateBalance(originAccountId string, destinationAccountId string, amount int) error {
	callInfo := struct {
		OriginAccountId      string
		DestinationAccountId string
		Amount               int
	}{
		OriginAccountId:      originAccountId,
		DestinationAccountId: destinationAccountId,
		Amount:               amount,
	}
	mock.lockUpdateBalance.Lock()
	mock.calls.UpdateBalance = append(mock.calls.UpdateBalance, callInfo)
	mock.lockUpdateBalance.Unlock()
	if mock.UpdateBalanceFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.UpdateBalanceFunc(originAccountId, destinationAccountId, amount)
}

// UpdateBalanceCalls gets all the calls that were made to UpdateBalance.
// Check the length with:
//     len(mockedUseCase.UpdateBalanceCalls())
func (mock *UseCaseMock) UpdateBalanceCalls() []struct {
	OriginAccountId      string
	DestinationAccountId string
	Amount               int
} {
	var calls []struct {
		OriginAccountId      string
		DestinationAccountId string
		Amount               int
	}
	mock.lockUpdateBalance.RLock()
	calls = mock.calls.UpdateBalance
	mock.lockUpdateBalance.RUnlock()
	return calls
}