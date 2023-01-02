package db

import (
	"context"
	"testing"

	"github.com/minhtam3010/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	acc := createRandomAccount(t)
	var amount int64
	getAmount := util.RandomMoney()

	if getAmount > acc.Balance {
		amount = getAmount - acc.Balance
	} else {
		amount = getAmount
	}
	rules := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    amount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), rules)

	require.NoError(t, err)
	require.Equal(t, rules.AccountID, entry.AccountID)
	require.Equal(t, rules.Amount, entry.Amount)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entryR := createRandomEntry(t)
	entry, err := testQueries.GetEntry(context.Background(), entryR.ID)

	require.NoError(t, err)
	require.Equal(t, entryR.AccountID, entry.AccountID)
	require.Equal(t, entryR.Amount, entry.Amount)
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	var amount int64
	getAmount := util.RandomMoney()

	if getAmount > entry.Amount {
		amount = getAmount - entry.Amount
	} else {
		amount = getAmount
	}

	rules := UpdateEntryParams{
		ID:     entry.ID,
		Amount: amount,
	}

	require.NoError(t, testQueries.UpdateEntry(context.Background(), rules))
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)

	require.NoError(t, testQueries.DeleteEntry(context.Background(), entry.ID))
}
