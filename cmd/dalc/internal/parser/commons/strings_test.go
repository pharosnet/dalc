package commons_test

import (
	"github.com/pharosnet/dalc/v2/cmd/dalc/internal/parser/commons"
	"testing"
)

func TestSnakeToCamel(t *testing.T) {
	s := "aggregateName"
	s = commons.SnakeToCamel(s)
	t.Log(s)
}
