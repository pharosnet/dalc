# dalc
Database access common layer for go.

## Feature

- Simple.
- No-reflect cost.
- Using callback function to decrease range times.
- Expandability.
- It is more convenient to use with dalg.

## Usage

`go get -u github.com/pharosnet/dalc`

```go
type SomeTableRow struct {

    // Fields mapped database columns
    // when using dalc.Scan(), then add col tag 
    Id string `col:"id"`

}

// insert func
// ctx := dal.WithPreparer(parentCtx, tx)
func Insert(ctx context.Context, rows... SomeTableInsert) (affected int64, err error) {

	affected, err = dalc.Execute(ctx, insertSql, 
		func(ctx context.Context, stmt *sql.Stmt, row interface{}) (result sql.Result, err error) {
		    someTableRow, ok := row.(*SomeTableRow)
		    if !ok {
		    	// 
		    }
		    result, err = stmt.ExecContext(ctx, someTableRow.Fields...)
		    return
	}, rows...)

	// other logic code

	return 
}

// ctx := dal.WithPreparer(parentCtx, tx or db)
func UseList() {

	someViews := make([]*SomeView , 0, 1)

	err = dalc.Query(ctx, ql, func(ctx context.Context, rows *sql.Rows, rowErr error) error {

			// check rowErr

			// rows -> SomeTableRow -> someView or use dalc.Scan(rows, &view)

			// someViews append someView
            return
	}, args...)

	// ...

}
```

## License

GNU GENERAL PUBLIC LICENSE 