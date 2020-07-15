package dalc

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// todo remove this file
func NowTime() *NullTime {
	return &NullTime{time.Now(), true}
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (n *NullTime) Scan(value interface{}) error {
	if value == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}
	switch value.(type) {
	case time.Time:
		n.Time = value.(time.Time)
		n.Valid = true
	}
	return nil
}

func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

func NewJson(v interface{}) (*NullJson, error) {
	p, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return &NullJson{Bytes: p, Valid: true}, nil
}

//type NullJson struct {
//	Bytes []byte `json:"-"`
//	Valid bool   `json:"-"`
//}

const tagName = "col"

type rowField struct {
	fieldIdx  int
	colIdx    int
	tag       string
	fieldType reflect.Type
}

func Scan(rows *sql.Rows, v interface{}) (err error) {

	if v == nil {
		err = errors.New("scan failed, cause target interface is nil")
		return
	}

	columnNames, colErr := rows.Columns()
	if colErr != nil {
		err = fmt.Errorf("scan failed, cause get columns from rows failed, %w", colErr)
		return
	}

	columns := make([]interface{}, 0, 1)
	for i := 0; i < len(columnNames); i++ {
		columns = append(columns, &sql.RawBytes{})
	}

	rt := reflect.TypeOf(v).Elem()
	rv := reflect.Indirect(reflect.ValueOf(v))

	rowFieldMappings := make([]*rowField, 0, 1)
	numOfField := rt.NumField()
	for i := 0; i < numOfField; i++ {
		field := rt.Field(i)
		tag, hasTag := field.Tag.Lookup(tagName)
		if !hasTag {
			continue
		}
		rowFieldMapping := &rowField{
			fieldIdx:  i,
			colIdx:    -1,
			tag:       tag,
			fieldType: field.Type,
		}

		for x, columnName := range columnNames {
			if strings.ToLower(columnName) != strings.ToLower(tag) {
				continue
			}
			v := reflect.New(field.Type).Interface()
			columns[x] = &v
			rowFieldMapping.colIdx = x

		}
		rowFieldMappings = append(rowFieldMappings, rowFieldMapping)
	}

	scanErr := rows.Scan(columns...)
	if scanErr != nil {
		panic(scanErr)
		return
	}

	for _, rowFieldMapping := range rowFieldMappings {
		if rowFieldMapping.colIdx < 0 {
			continue
		}
		cv := reflect.Indirect(reflect.ValueOf(columns[rowFieldMapping.colIdx]))
		reflect.Indirect(rv.Field(rowFieldMapping.fieldIdx)).Set(cv.Elem())
	}

	return
}
