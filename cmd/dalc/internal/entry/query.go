package entry

const (
	ReadQueryKind  = QueryKind("r")
	WriteQueryKind = QueryKind("w")
)

type QueryKind string

type Query struct {
	Kind   QueryKind
	Name   string
	Ref    *Table
	Result *QueryResult
}

type QueryResult struct {
	Fields []*QueryResultField
}

type QueryResultField struct {
	Name string
	Type string
}
