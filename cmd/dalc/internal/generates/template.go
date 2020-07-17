package generates

import "github.com/pharosnet/dalc/cmd/dalc/internal/entry"

type TemplateData struct {
	Package string
	Imports map[string]string
}

type GenerateTableData struct {
	RawName       string
	Package       string
	Imports       map[string]string
	LowName       string
	Name          string
	GetSQL        string
	InsertSQL     string
	UpdateSQL     string
	DeleteSQL     string
	HasAutoIncrId bool
	Fields        []*TableField
}

type TableField struct {
	Pk       bool
	AutoIncr bool
	Name     string
	Type     string
	Tags     string
}

type QueryField struct {
	Name string
	Type string
	Tags string
}

type GenerateQueryData struct {
	Exec          bool
	RawName       string
	Package       string
	Imports       map[string]string
	LowName       string
	Name          string
	QuerySQL      string
	RequestFields []*QueryField
	IsTable       bool
	Table         *entry.Table
	ResultFields  []*QueryField
}
