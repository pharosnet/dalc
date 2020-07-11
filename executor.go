package dalc

import (
	"context"
	"errors"
	"fmt"
)

func Execute(ctx context.Context, query string, args *Args) (affected int64, err error) {
	if ctx == nil {
		err = errors.New("execute failed, context is empty")
		return
	}
	if query == "" {
		err = errors.New("execute failed, query is empty")
		return
	}
	if args == nil || args.IsEmpty() {
		err = errors.New("execute failed, args are empty")
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
	result, execErr := stmt.ExecContext(ctx, args.values)
	if execErr != nil {
		err = execErr
		return
	}
	affectedRows, affectedErr := result.RowsAffected()
	if affectedErr != nil {
		err = fmt.Errorf("execute failed, get result's affected failed. sql: %s reason: %v", query, affectedErr)
		return
	}
	if affectedRows == 0 {
		err = fmt.Errorf("execute failed, affected nothing, sql: %s", query)
		return
	}
	affected = affectedRows
	if hasLog() {
		logger.Debugf("execute success, affected: %d, sql: %s", affected, query)
	}
	return
}
