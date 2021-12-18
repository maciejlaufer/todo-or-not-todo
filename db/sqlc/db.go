// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createListStmt, err = db.PrepareContext(ctx, createList); err != nil {
		return nil, fmt.Errorf("error preparing query CreateList: %w", err)
	}
	if q.createTaskStmt, err = db.PrepareContext(ctx, createTask); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTask: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteListStmt, err = db.PrepareContext(ctx, deleteList); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteList: %w", err)
	}
	if q.deleteTaskStmt, err = db.PrepareContext(ctx, deleteTask); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTask: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getListStmt, err = db.PrepareContext(ctx, getList); err != nil {
		return nil, fmt.Errorf("error preparing query GetList: %w", err)
	}
	if q.getListsByCreatorIdStmt, err = db.PrepareContext(ctx, getListsByCreatorId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListsByCreatorId: %w", err)
	}
	if q.getTaskStmt, err = db.PrepareContext(ctx, getTask); err != nil {
		return nil, fmt.Errorf("error preparing query GetTask: %w", err)
	}
	if q.getTasksByListIdStmt, err = db.PrepareContext(ctx, getTasksByListId); err != nil {
		return nil, fmt.Errorf("error preparing query GetTasksByListId: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.updateListStmt, err = db.PrepareContext(ctx, updateList); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateList: %w", err)
	}
	if q.updateTaskStmt, err = db.PrepareContext(ctx, updateTask); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTask: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createListStmt != nil {
		if cerr := q.createListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createListStmt: %w", cerr)
		}
	}
	if q.createTaskStmt != nil {
		if cerr := q.createTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTaskStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteListStmt != nil {
		if cerr := q.deleteListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteListStmt: %w", cerr)
		}
	}
	if q.deleteTaskStmt != nil {
		if cerr := q.deleteTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTaskStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getListStmt != nil {
		if cerr := q.getListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListStmt: %w", cerr)
		}
	}
	if q.getListsByCreatorIdStmt != nil {
		if cerr := q.getListsByCreatorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListsByCreatorIdStmt: %w", cerr)
		}
	}
	if q.getTaskStmt != nil {
		if cerr := q.getTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTaskStmt: %w", cerr)
		}
	}
	if q.getTasksByListIdStmt != nil {
		if cerr := q.getTasksByListIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTasksByListIdStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.updateListStmt != nil {
		if cerr := q.updateListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateListStmt: %w", cerr)
		}
	}
	if q.updateTaskStmt != nil {
		if cerr := q.updateTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTaskStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                      DBTX
	tx                      *sql.Tx
	createListStmt          *sql.Stmt
	createTaskStmt          *sql.Stmt
	createUserStmt          *sql.Stmt
	deleteListStmt          *sql.Stmt
	deleteTaskStmt          *sql.Stmt
	deleteUserStmt          *sql.Stmt
	getListStmt             *sql.Stmt
	getListsByCreatorIdStmt *sql.Stmt
	getTaskStmt             *sql.Stmt
	getTasksByListIdStmt    *sql.Stmt
	getUserByEmailStmt      *sql.Stmt
	getUserByIdStmt         *sql.Stmt
	getUsersStmt            *sql.Stmt
	updateListStmt          *sql.Stmt
	updateTaskStmt          *sql.Stmt
	updateUserStmt          *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                      tx,
		tx:                      tx,
		createListStmt:          q.createListStmt,
		createTaskStmt:          q.createTaskStmt,
		createUserStmt:          q.createUserStmt,
		deleteListStmt:          q.deleteListStmt,
		deleteTaskStmt:          q.deleteTaskStmt,
		deleteUserStmt:          q.deleteUserStmt,
		getListStmt:             q.getListStmt,
		getListsByCreatorIdStmt: q.getListsByCreatorIdStmt,
		getTaskStmt:             q.getTaskStmt,
		getTasksByListIdStmt:    q.getTasksByListIdStmt,
		getUserByEmailStmt:      q.getUserByEmailStmt,
		getUserByIdStmt:         q.getUserByIdStmt,
		getUsersStmt:            q.getUsersStmt,
		updateListStmt:          q.updateListStmt,
		updateTaskStmt:          q.updateTaskStmt,
		updateUserStmt:          q.updateUserStmt,
	}
}
