package entry

import (
	"fmt"
)

type ColumnType string

func (ct ColumnType) GoType() (goType *GoType, err error) {

	switch ct {
	case TINYINT, SMALLINT, MEDIUMINT, INT, INTEGER:
		goType = NewGoType("sql.NullInt32")
	case BIGINT:
		goType = NewGoType("sql.NullInt64")
	case FLOAT, DOUBLE, DECIMAL:
		goType = NewGoType("sql.NullFloat64")
	case BOOLEAN:
		goType = NewGoType("sql.NullBool")
	case DATE, TIME, YEAR, DATETIME, TIMESTAMP:
		goType = NewGoType("sql.NullTime")
	case CHAR, NCHAR, VARCHAR, NVARCHAR:
		goType = NewGoType("sql.NullString")
	case TINYBLOB, TINYTEXT, BLOB, TEXT, MEDIUMBLOB, MEDIUMTEXT, LONGBLOB, LONGTEXT:
		goType = NewGoType("[]byte")
	case JSON:
		goType = NewGoType("json.RawMessage")
	default:
		err = fmt.Errorf("column type mapped no go type, %v", ct)
	}

	return
}

type Column struct {
	Name         string
	Type         ColumnType
	GoType       *GoType
	DefaultValue string
}
