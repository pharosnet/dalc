package dalc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type QueryScanner interface {
	Scanner(ctx context.Context, rows *sql.Rows) (err error)
	OnRowError(err error) (breakOff bool)
}

func Query(ctx context.Context, query string, scanner QueryScanner, args ...interface{}) (err error) {
	if ctx == nil {
		err = errors.New("dalc-> query failed, context is empty")
		return
	}
	if query == "" {
		err = errors.New("dalc-> query failed, query is empty")
		return
	}
	if scanner == nil {
		err = errors.New("dalc-> query failed, scanner is empty")
		return
	}
	stmt, prepareErr := prepare(ctx).PrepareContext(ctx, query)
	if prepareErr != nil {
		err = fmt.Errorf("dalc-> query failed, prepared statement failed. reason: %v", prepareErr)
		return
	}
	defer func() {
		stmtCloseErr := stmt.Close()
		if stmtCloseErr != nil {
			err = fmt.Errorf("dalc-> query failed, close prepare statement failed. reason: %v", stmtCloseErr)
			return
		}
	}()
	rows, queryErr := stmt.QueryContext(ctx, args...)
	if queryErr != nil {
		err = fmt.Errorf("dalc-> query failed, query failed. reason: %v", queryErr)
		return
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			err = fmt.Errorf("dalc-> query failed, close rows failed. reason: %v", closeErr)
			return
		}
	}()
	for rows.Next() {
		if rowErr := rows.Err(); rowErr != nil {
			if breakOff := scanner.OnRowError(rowErr); breakOff {
				err = rowErr
				return err
			}
		}
		scanErr := scanner.Scanner(ctx, rows)
		if scanErr != nil {
			err = scanErr
			return
		}
	}
	if hasLog() {
		logf("dalc-> query success, sql: %s\n", query)
	}
	return
}
