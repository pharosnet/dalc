package dalc

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

func MySqlTcpDSN(name string, password string, address string, schema string, parseTime bool, loc string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=%v&loc=%s",
		name, password, address, schema, parseTime, loc,
	)
}

// name:password@tcp(ip:host)/schema?parseTime=true&loc=Local
type MySQLTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (n *MySQLTime) Scan(value interface{}) (err error) {
	if value == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}
	n.Valid = true
	switch value.(type) {
	case []uint8:
		v, _ := value.([]uint8)
		n.Time, err = time.ParseInLocation("2006-01-02 15:04:05", string(v), time.Local)
	case string:
		v, _ := value.(string)
		n.Time, err = time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
	case time.Time:
		n.Time = value.(time.Time)
	default:
		err = fmt.Errorf("dalc scan mysql time type failed, %s is not []uint8 and string", reflect.TypeOf(value).Name())
	}
	return
}

// Value implements the driver Valuer interface.
func (n MySQLTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
