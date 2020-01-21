package dalc_test

import (
	"testing"

	"github.com/pharosnet/dalc"
)

func TestNewVarchar(t *testing.T) {
	v := dalc.NewVarchar("sss")
	t.Log(v)
}
