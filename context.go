package dalc

import (
	"context"
	"database/sql"
	"errors"
)

const (
	ctxKeyPreparedStatement = "_dalc_prepared_statement"
)

type PreparedStatement interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func WithPreparedStatement(parent context.Context, p PreparedStatement) PreparedContext {
	return context.WithValue(parent, ctxKeyPreparedStatement, p)
}

func prepare(ctx PreparedContext) PreparedStatement {
	v := ctx.Value(ctxKeyPreparedStatement)
	if v == nil {
		panic(errors.New("the ctx is not dalc.PreparedContext"))
		return nil
	}
	return v.(PreparedStatement)
}

type PreparedContext context.Context
