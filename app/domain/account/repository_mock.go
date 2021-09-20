// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package account

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
// 			CreateAccountFunc: func(ctx context.Context, account entities.Account) error {
// 				panic("mock out the CreateAccount method")
// 			},
// 			GetAccountsFunc: func(ctx context.Context) ([]entities.Account, error) {
// 				panic("mock out the GetAccounts method")
// 			},
// 			GetByCpfFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
// 				panic("mock out the GetByCpf method")
// 			},
// 			GetByIdFunc: func(ctx context.Context, id string) (entities.Account, error) {
// 				panic("mock out the GetById method")
// 			},
// 			UpdateBalanceFunc: func(ctx context.Context, id string, value int) error {
// 				panic("mock out the UpdateBalance method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// CreateAccountFunc mocks the CreateAccount method.
	CreateAccountFunc func(ctx context.Context, account entities.Account) error

	// GetAccountsFunc mocks the GetAccounts method.
	GetAccountsFunc func(ctx context.Context) ([]entities.Account, error)

	// GetByCpfFunc mocks the GetByCpf method.
	GetByCpfFunc func(ctx context.Context, cpf string) (entities.Account, error)

	// GetByIdFunc mocks the GetById method.
	GetByIdFunc func(ctx context.Context, id string) (entities.Account, error)

	// UpdateBalanceFunc mocks the UpdateBalance method.
	UpdateBalanceFunc func(ctx context.Context, id string, value int) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateAccount holds details about calls to the CreateAccount method.
		CreateAccount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Account is the account argument value.
			Account entities.Account
		}
		// GetAccounts holds details about calls to the GetAccounts method.
		GetAccounts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetByCpf holds details about calls to the GetByCpf method.
		GetByCpf []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Cpf is the cpf argument value.
			Cpf string
		}
		// GetById holds details about calls to the GetById method.
		GetById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// UpdateBalance holds details about calls to the UpdateBalance method.
		UpdateBalance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
			// Value is the value argument value.
			Value int
		}
	}
	lockCreateAccount sync.RWMutex
	lockGetAccounts   sync.RWMutex
	lockGetByCpf      sync.RWMutex
	lockGetById       sync.RWMutex
	lockUpdateBalance sync.RWMutex
}

// CreateAccount calls CreateAccountFunc.
func (mock *RepositoryMock) CreateAccount(ctx context.Context, account entities.Account) error {
	callInfo := struct {
		Ctx     context.Context
		Account entities.Account
	}{
		Ctx:     ctx,
		Account: account,
	}
	mock.lockCreateAccount.Lock()
	mock.calls.CreateAccount = append(mock.calls.CreateAccount, callInfo)
	mock.lockCreateAccount.Unlock()
	if mock.CreateAccountFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.CreateAccountFunc(ctx, account)
}

// CreateAccountCalls gets all the calls that were made to CreateAccount.
// Check the length with:
//     len(mockedRepository.CreateAccountCalls())
func (mock *RepositoryMock) CreateAccountCalls() []struct {
	Ctx     context.Context
	Account entities.Account
} {
	var calls []struct {
		Ctx     context.Context
		Account entities.Account
	}
	mock.lockCreateAccount.RLock()
	calls = mock.calls.CreateAccount
	mock.lockCreateAccount.RUnlock()
	return calls
}

// GetAccounts calls GetAccountsFunc.
func (mock *RepositoryMock) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
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
	return mock.GetAccountsFunc(ctx)
}

// GetAccountsCalls gets all the calls that were made to GetAccounts.
// Check the length with:
//     len(mockedRepository.GetAccountsCalls())
func (mock *RepositoryMock) GetAccountsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAccounts.RLock()
	calls = mock.calls.GetAccounts
	mock.lockGetAccounts.RUnlock()
	return calls
}

// GetByCpf calls GetByCpfFunc.
func (mock *RepositoryMock) GetByCpf(ctx context.Context, cpf string) (entities.Account, error) {
	callInfo := struct {
		Ctx context.Context
		Cpf string
	}{
		Ctx: ctx,
		Cpf: cpf,
	}
	mock.lockGetByCpf.Lock()
	mock.calls.GetByCpf = append(mock.calls.GetByCpf, callInfo)
	mock.lockGetByCpf.Unlock()
	if mock.GetByCpfFunc == nil {
		var (
			accountOut entities.Account
			errOut     error
		)
		return accountOut, errOut
	}
	return mock.GetByCpfFunc(ctx, cpf)
}

// GetByCpfCalls gets all the calls that were made to GetByCpf.
// Check the length with:
//     len(mockedRepository.GetByCpfCalls())
func (mock *RepositoryMock) GetByCpfCalls() []struct {
	Ctx context.Context
	Cpf string
} {
	var calls []struct {
		Ctx context.Context
		Cpf string
	}
	mock.lockGetByCpf.RLock()
	calls = mock.calls.GetByCpf
	mock.lockGetByCpf.RUnlock()
	return calls
}

// GetById calls GetByIdFunc.
func (mock *RepositoryMock) GetById(ctx context.Context, id string) (entities.Account, error) {
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
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
	return mock.GetByIdFunc(ctx, id)
}

// GetByIdCalls gets all the calls that were made to GetById.
// Check the length with:
//     len(mockedRepository.GetByIdCalls())
func (mock *RepositoryMock) GetByIdCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetById.RLock()
	calls = mock.calls.GetById
	mock.lockGetById.RUnlock()
	return calls
}

// UpdateBalance calls UpdateBalanceFunc.
func (mock *RepositoryMock) UpdateBalance(ctx context.Context, id string, value int) error {
	callInfo := struct {
		Ctx   context.Context
		ID    string
		Value int
	}{
		Ctx:   ctx,
		ID:    id,
		Value: value,
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
	return mock.UpdateBalanceFunc(ctx, id, value)
}

// UpdateBalanceCalls gets all the calls that were made to UpdateBalance.
// Check the length with:
//     len(mockedRepository.UpdateBalanceCalls())
func (mock *RepositoryMock) UpdateBalanceCalls() []struct {
	Ctx   context.Context
	ID    string
	Value int
} {
	var calls []struct {
		Ctx   context.Context
		ID    string
		Value int
	}
	mock.lockUpdateBalance.RLock()
	calls = mock.calls.UpdateBalance
	mock.lockUpdateBalance.RUnlock()
	return calls
}