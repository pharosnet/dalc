package dalc

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

var emptyArgs = &Args{
	values: make([]interface{}, 0, 1),
}

func NewArgs() *Args {
	return &Args{
		values: make([]interface{}, 0, 1),
	}
}

type Args struct {
	values []interface{}
}

func (a *Args) Arg(v interface{}) *Args {
	a.values = append(a.values, v)
	return a
}

func (a *Args) Merge(args *Args) *Args {
	a.values = append(a.values, args.values...)
	return a
}

func (a *Args) Values() []interface{} {
	return a.values
}

func (a *Args) IsEmpty() bool {
	return len(a.values) == 0
}

func (a Args) String() string {
	return fmt.Sprintf("%v", a.values)
}

func NewTupleArgs(v interface{}) (args TupleArgs) {
	rt := reflect.ValueOf(v)
	num := rt.Len()
	args = make([]interface{}, 0, 1)
	for i := 0; i < num; i++ {
		args = append(args, rt.Index(i).Interface())
	}
	return
}

type TupleArgs []interface{}

func ReplaceSQL(query string, key string, args TupleArgs) string {
	buf := bytes.NewBufferString("")
	for i, arg := range args {
		if i > 0 {
			buf.WriteString(", ")
		}
		switch arg.(type) {
		case string:
			buf.WriteString(fmt.Sprintf("'%s'", arg.(string)))
		case []byte:
			buf.WriteString(fmt.Sprintf("'%s'", string(arg.([]byte))))
		case int:
			buf.WriteString(fmt.Sprintf("%d", arg.(int)))
		case int32:
			buf.WriteString(fmt.Sprintf("%d", arg.(int32)))
		case int64:
			buf.WriteString(fmt.Sprintf("%d", arg.(int64)))
		case float32:
			buf.WriteString(fmt.Sprintf("%d", arg.(float32)))
		case float64:
			buf.WriteString(fmt.Sprintf("%d", arg.(float64)))
		default:
			panic(fmt.Errorf("dalc replace sql failed, unsupport arg type, %v, %s", reflect.TypeOf(arg), query))
		}
	}
	return strings.Replace(query, fmt.Sprintf("#%s#", key), buf.String(), 1)
}
