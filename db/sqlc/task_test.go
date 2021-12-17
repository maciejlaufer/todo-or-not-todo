package db

import (
	"context"
	"database/sql"
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

func TestDeleteTask(t *testing.T) {
	task1 := createRandomTask(t)
	err := testQueries.DeleteTask(context.Background(), task1.ID)
	require.NoError(t, err)

	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task2)
}

func TestUpdateTask(t *testing.T) {
	task1 := createRandomTask(t)

	arg := UpdateTaskParams{
		ID:          task1.ID,
		Name:        util.RandomString(6),
		Description: NewNullString(util.RandomString(16)),
	}

	task2, err := testQueries.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, arg.Name, task2.Name)
	require.Equal(t, arg.Description.String, task2.Description.String)
	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second)
}

func TestGetTasks(t *testing.T) {
	task := createRandomTask(t)

	tasks, err := testQueries.GetTasksByListId(context.Background(), task.ListID)
	require.NoError(t, err)
	require.Len(t, tasks, 1)

	for _, task := range tasks {
		require.NotEmpty(t, task)
	}
}
