package dalc

import (
	"context"
	"database/sql"
)

const (
	ctxKeyPreparable = "_preparable"
)

type Preparable interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func WithPreparable(parent context.Context, p Preparable) context.Context {
	return context.WithValue(parent, ctxKeyPreparable, p)
}

func prepare(ctx context.Context) Preparable {
	v := ctx.Value(ctxKeyPreparable)
	if v == nil {
		return nil
	}
	return v.(Preparable)
}
