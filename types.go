package dalc

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type NullBytes struct {
	Bytes []byte
	Valid bool // Valid is true if Time is not NULL
}

func (n *NullBytes) Scan(value interface{}) (err error) {
	if value == nil {
		n.Bytes, n.Valid = nil, false
		return nil
	}
	n.Valid = true
	switch value.(type) {
	case []byte:
		n.Bytes = value.([]byte)
	case string:
		n.Bytes = []byte(value.(string))
	default:
		err = fmt.Errorf("dalc scan mysql time type failed, %s is not []uint8 and string", reflect.TypeOf(value).Name())
	}
	return
}

// Value implements the driver Valuer interface.
func (n NullBytes) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bytes, nil
}

type NullJson NullBytes

func (n *NullJson) Unmarshal(v interface{}) (err error) {
	if !n.Valid {
		return
	}
	err = json.Unmarshal(n.Bytes, v)
	return
}

func (n *NullJson) Marshal(v interface{}) (err error) {

	if reflect.ValueOf(v).IsNil() {
		n.Valid = false
		return
	}

	p, err0 := json.Marshal(v)
	if err0 != nil {
		n.Valid = false
		err = err0
		return
	}

	n.Valid = true
	n.Bytes = p
	return
}
