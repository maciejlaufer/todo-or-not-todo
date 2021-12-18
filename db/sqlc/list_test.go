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
		Name:      util.RandomString(10),
		CreatorID: user.ID,
	}

	list, err := testQueries.CreateList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	require.Equal(t, arg.Name, list.Name)
	require.Equal(t, arg.CreatorID, list.CreatorID)

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

func TestUpdateList(t *testing.T) {
	list1 := createRandomList(t)

	arg := UpdateListParams{
		ID:   list1.ID,
		Name: util.RandomString(6),
	}

	list2, err := testQueries.UpdateList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list2)

	require.Equal(t, list1.ID, list2.ID)
	require.Equal(t, arg.Name, list2.Name)
	require.WithinDuration(t, list1.CreatedAt, list2.CreatedAt, time.Second)
}

func TestGetLists(t *testing.T) {
	list := createRandomList(t)

	lists, err := testQueries.GetListsByCreatorId(context.Background(), list.CreatorID)
	require.NoError(t, err)
	require.Len(t, lists, 1)

	for _, list := range lists {
		require.NotEmpty(t, list)
	}
}
