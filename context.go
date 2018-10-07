package dalc

import (
	"context"
	"database/sql"
)

const (
	ctxKeyPreparer = "preparer"
)

type Preparer interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func WithPreparer(parent context.Context, p Preparer) context.Context {
	return context.WithValue(parent, ctxKeyPreparer, p)
}

func prepare(ctx context.Context) Preparer {
	v := ctx.Value(ctxKeyPreparer)
	if v == nil {
		return nil
	}
	return v.(Preparer)
}
