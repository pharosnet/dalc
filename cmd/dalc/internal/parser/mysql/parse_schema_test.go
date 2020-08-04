package mysql_test

import (
	"github.com/pharosnet/dalc/v2/cmd/dalc/v2/internal/parser/mysql"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestParseMySQLSchema(t *testing.T) {

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		t.Error(pwdErr)
		return
	}

	schemaPath := filepath.Join(pwd, "schema.sql")

	p, pErr := ioutil.ReadFile(schemaPath)
	if pErr != nil {
		t.Error(pErr)
		return
	}

	schema, parseErr := mysql.ParseMySQLSchema(string(p))
	if parseErr != nil {
		t.Error(parseErr)
		return
	}
	t.Log(schema.Name, len(schema.Tables))
	for _, table := range schema.Tables {
		t.Log("table", table.FullName, table.Schema, table.Name, table.GoName)
		t.Log("\t", "pk", table.PKs)
		for _, column := range table.Columns {
			t.Log("\t", "column", column.GoName, column.Name, column.Type, column.DefaultValue, column.GoType, column.Null, column.AutoIncrement)
		}
	}
}
