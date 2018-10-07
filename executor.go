package dalc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Executable interface {
	Exec(ctx context.Context, stmt *sql.Stmt) (result sql.Result, err error)
}

func Execute(ctx context.Context, query string, executables ...Executable) (affected int64, err error) {
	if ctx == nil {
		err = errors.New("dalc-> execute failed, context is empty")
		return
	}
	if query == "" {
		err = errors.New("dalc-> execute failed, query is empty")
		return
	}
	if executables == nil || len(executables) == 0 {
		err = errors.New("dalc-> execute failed, executables are empty")
		return
	}
	stmt, prepareErr := prepare(ctx).PrepareContext(ctx, query)
	if prepareErr != nil {
		err = fmt.Errorf("dalc-> execute failed, prepared statement failed. reason: %v", prepareErr)
		return
	}
	defer func() {
		stmtCloseErr := stmt.Close()
		if stmtCloseErr != nil {
			err = fmt.Errorf("dalc-> execute failed, close prepare statement failed. reason: %v", stmtCloseErr)
			return
		}
	}()
	for _, executable := range executables {
		result, execErr := executable.Exec(ctx, stmt)
		if execErr != nil {
			err = execErr
			return
		}
		affectedRows, affectedErr := result.RowsAffected()
		if affectedErr != nil {
			err = fmt.Errorf("dalc-> execute failed, get result's affected failed. reason: %v, executable: %v", affectedErr, executable)
			return
		}
		if affectedRows == 0 {
			err = fmt.Errorf("dalc-> execute failed, affected nothing, executable: %v", executables)
			return
		}
		affected = affected + affectedRows
	}
	if hasLog() {
		logf("dalc-> execute success, affected: %d, sql: %s", affected, query)
	}
	return
}
