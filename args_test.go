package dalc_test

import (
	"github.com/pharosnet/dalc/v2"
	"testing"
)

func TestNewArgs(t *testing.T) {
	args := dalc.NewArgs()
	argsFunc(t, args.Arg("1").Arg("2"))
	argsFunc(t)
}

func argsFunc(t *testing.T, args ...interface{}) {
	t.Log(args == nil, len(args), args)

}
