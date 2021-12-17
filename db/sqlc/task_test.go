package db

import (
	"context"
	"testing"
	"time"

	"github.com/maciejlaufer/todoornottodo/util"
	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) Task {
	list := createRandomList(t)

	arg := CreateTaskParams{
		Name:        util.RandomString(10),
		Description: NewNullString(util.RandomString(20)),
		ListID:      list.ID,
	}

	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	require.Equal(t, arg.Name, task.Name)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.ListID, task.ListID)

	require.NotZero(t, task.ID)
	require.NotZero(t, task.CreatedAt)

	return task
}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := createRandomTask(t)
	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, task1.Name, task2.Name)
	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second)
}
