package entry

type Table struct {
	FullName string
	Schema   string
	Name     string
	Columns  []*Column
}
