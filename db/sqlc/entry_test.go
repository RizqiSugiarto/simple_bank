package db

import (
	"context"
	"database/sql"
	"learn-until-die/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	args := CreateEntryParams{
		AccountID: util.RandomId(),
		Amount:    util.RandomMoney(),
	}
	Entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, Entry)

	require.Equal(t, args.AccountID, Entry.AccountID)
	require.Equal(t, args.Amount, Entry.Amount)

	require.NotZero(t, Entry.ID)
	require.NotZero(t, Entry.CreatedAt)

	return Entry

}

func TestCreateRandomEntry(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestUpdateEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	args := UpdateEntryParams{
		ID:        entry1.ID,
		AccountID: 1,
		Amount:    util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), args)

	require.NoError(t, err)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, args.AccountID, entry2.AccountID)
	require.Equal(t, args.Amount, args.Amount)

}

func TestListEntry(t *testing.T) {
	args := ListEntryParams{
		Limit:  5,
		Offset: 5,
	}

	entrys, err := testQueries.ListEntry(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, entrys, 5)

	for _, entry := range entrys {
		require.NotEmpty(t, entry)
	}
}

func TestDeleteEntry(t *testing.T) {
	err := testQueries.DeleteEntry(context.Background(), 18)

	require.NoError(t, err)

	entrys, err := testQueries.GetEntry(context.Background(), 18)

	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, entrys)
}
