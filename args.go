package dalc

import "fmt"

var emptyArgs = &Args{
	values: make([]interface{}, 0, 1),
}

func NewArgs() Args {
	return Args{
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

func (a *Args) Values() []interface{} {
	return a.values
}

func (a *Args) IsEmpty() bool {
	return len(a.values) == 0
}

func (a Args) String() string {
	return fmt.Sprintf("%v", a.values)
}
