package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/maciejlaufer/todoornottodo/util"
	"github.com/stretchr/testify/require"
)

func createRandomList(t *testing.T) List {
	user := createRandomUser(t)

	arg := CreateListParams{
		Name:   util.RandomString(10),
		UserID: user.ID,
	}

	list, err := testQueries.CreateList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	require.Equal(t, arg.Name, list.Name)
	require.Equal(t, arg.UserID, list.UserID)

	require.NotZero(t, list.ID)
	require.NotZero(t, list.CreatedAt)

	return list
}

func TestCreateList(t *testing.T) {
	createRandomList(t)
}

func TestGetList(t *testing.T) {
	list1 := createRandomList(t)
	list2, err := testQueries.GetList(context.Background(), list1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, list2)

	require.Equal(t, list1.ID, list2.ID)
	require.Equal(t, list1.Name, list2.Name)
	require.WithinDuration(t, list1.CreatedAt, list2.CreatedAt, time.Second)
}

func TestDeleteList(t *testing.T) {
	list1 := createRandomList(t)
	err := testQueries.DeleteList(context.Background(), list1.ID)
	require.NoError(t, err)

	list2, err := testQueries.GetList(context.Background(), list1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, list2)
}
