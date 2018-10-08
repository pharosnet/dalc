package dalc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type ExecuteFunc func(ctx context.Context, stmt *sql.Stmt, row interface{}) (result sql.Result, err error)

func Execute(ctx context.Context, query string, fn ExecuteFunc, rows ...interface{}) (affected int64, err error) {
	if ctx == nil {
		err = errors.New("execute failed, context is empty")
		return
	}
	if query == "" {
		err = errors.New("execute failed, query is empty")
		return
	}
	if fn == nil {
		err = errors.New("execute failed, execute function is empty")
		return
	}
	if rows == nil || len(rows) == 0 {
		err = errors.New("execute failed, rows are empty")
		return
	}
	stmt, prepareErr := prepare(ctx).PrepareContext(ctx, query)
	if prepareErr != nil {
		err = fmt.Errorf("execute failed, prepared statement failed. reason: %v", prepareErr)
		return
	}
	defer func() {
		stmtCloseErr := stmt.Close()
		if stmtCloseErr != nil {
			err = fmt.Errorf("execute failed, close prepare statement failed. reason: %v", stmtCloseErr)
			return
		}
	}()
	for _, row := range rows {
		result, execErr := fn(ctx, stmt, row)
		if execErr != nil {
			err = execErr
			return
		}
		affectedRows, affectedErr := result.RowsAffected()
		if affectedErr != nil {
			err = fmt.Errorf("execute failed, get result's affected failed. reason: %v, row: %v", affectedErr, row)
			return
		}
		if affectedRows == 0 {
			err = fmt.Errorf("execute failed, affected nothing, row: %v", row)
			return
		}
		affected = affected + affectedRows
	}
	if hasLog() {
		logf("execute success, affected: %d, sql: %s", affected, query)
	}
	return
}
