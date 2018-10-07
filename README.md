# dalc
Database access common layer for go.

## Feature

- Simple.
- No-reflect cost.
- Using callback function to decrease range times.
- expandability.
- It is more convenient to use with dalg.

## Usage

`go get -u github.com/pharosnet/dalc`

```go
type SomeTableRow struct {

// Fields mapped database columns

}

// insert definition

type SomeTableInsert struct {

	SomeTableRow 

}

func (t *SomeTableInsert) Exec(ctx context.Context, stmt *sql.Stmt) (result sql.Result, err error) {

	result, err = stmt.ExecContext(ctx, t.Fields...)

	// other logic code

	return

}

// insert func
// ctx := dal.WithPreparer(parentCtx, tx)

func Insert(ctx context.Context, rows... SomeTableInsert) (affected int64, err error) {

	affected, err := dalc.Execute(ctx, insertSql, rows...)

	// other logic code

	return 

}

// query definition 

func List(ctx context.Context, rangeFn SomeTableRowScanner) (err error) {

	err = dalc.Query(ctx, querySql, rangeFn, args...)

	return

}

// ctx := dal.WithPreparer(parentCtx, tx or db)
func UseList() {

	someViews := make([]*SomeView , 0, 1)

	err = List(ctx, func(ctx context.Context, rows *sql.Rows, rowErr error) error {

			// check rowErr

			// rows -> SomeTableRow -> someView

			// someViews append someView

	})

	// ...

}
```

## License

GNU GENERAL PUBLIC LICENSE 