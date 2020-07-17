package dalc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type QueryResultIterator func(ctx context.Context, rows *sql.Rows, rowErr error) (err error)

func Query(ctx PreparedContext, query string, args *Args, iterator QueryResultIterator) (err error) {
	if ctx == nil {
		err = errors.New("query failed, context is empty")
		return
	}
	if query == "" {
		err = errors.New("query failed, query is empty")
		return
	}
	if iterator == nil {
		err = errors.New("query failed, iterator is empty")
		return
	}
	stmt, prepareErr := prepare(ctx).PrepareContext(ctx, query)
	if prepareErr != nil {
		err = fmt.Errorf("query failed, prepared statement failed. reason: %v", prepareErr)
		return
	}
	defer func() {
		stmtCloseErr := stmt.Close()
		if stmtCloseErr != nil {
			err = fmt.Errorf("query failed, close prepare statement failed. reason: %v", stmtCloseErr)
			return
		}
	}()
	var rows *sql.Rows = nil
	var queryErr error = nil
	if args != nil && len(args.values) > 0 {
		rows, queryErr = stmt.QueryContext(ctx, args.Values()...)
	} else {
		rows, queryErr = stmt.QueryContext(ctx)
	}

	if queryErr != nil {
		err = fmt.Errorf("query failed, query failed. reason: %v", queryErr)
		return
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			err = fmt.Errorf("query failed, close rows failed. reason: %v", closeErr)
			return
		}
	}()
	for rows.Next() {
		iteratorErr := iterator(ctx, rows, rows.Err())
		if iteratorErr != nil {
			err = iteratorErr
			return
		}
	}
	if hasLog() {
		logger.Debugf("query success, sql: %s\n", query)
	}
	return
}
