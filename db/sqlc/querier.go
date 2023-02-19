// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	AddUserBalance(ctx context.Context, arg AddUserBalanceParams) (UserAccount, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (UserAccount, error)
	DeleteUser(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (UserAccount, error)
	GetAccountForUpdate(ctx context.Context, id int64) (UserAccount, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]UserAccount, error)
	SubUserBalance(ctx context.Context, arg SubUserBalanceParams) (UserAccount, error)
}

var _ Querier = (*Queries)(nil)