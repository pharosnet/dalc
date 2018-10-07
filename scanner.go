package dalc

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

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

type NullJson struct {
	Bytes []byte `json:"-"`
	Valid bool   `json:"-"`
}
