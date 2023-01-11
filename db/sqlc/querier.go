// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	DeleteTransfers(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTranfers(ctx context.Context, id int64) (Transfer, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) error
	UpdateTransfers(ctx context.Context, arg UpdateTransfersParams) error
}

var _ Querier = (*Queries)(nil)
