package db

import (
	"context"
	"testing"

	"github.com/minhtam3010/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	rules := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), rules)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, rules.Owner, account.Owner)
	require.Equal(t, rules.Currency, account.Currency)
	require.Equal(t, rules.Balance, account.Balance)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountR := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), accountR.ID)

	require.NoError(t, err)
	require.Equal(t, accountR.Owner, account.Owner)
	require.Equal(t, accountR.Balance, account.Balance)
	require.Equal(t, accountR.Currency, account.Currency)
}

// func TestListAccount(t *testing.T) {
// 	limit := ListAccountsParams{
// 		Limit:  5,
// 		Offset: 0,
// 	}
// 	account, err := testQueries.ListAccounts(context.Background(), limit)
// 	require.NoError(t, err)
// 	require.Equal(t, 5, len(account))
// }

func TestUpdateAccount(t *testing.T) {
	acc := createRandomAccount(t)

	updateAcc := UpdateAccountParams{
		ID:      acc.ID,
		Balance: util.RandomMoney(),
	}

	err := testQueries.UpdateAccount(context.Background(), updateAcc)
	require.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)
}
