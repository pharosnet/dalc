package commons_test

import (
	"github.com/pharosnet/dalc/cmd/dalc/v2/internal/parser/commons"
	"testing"
)

func TestSnakeToCamel(t *testing.T) {
	s := "aggregateName"
	s = commons.SnakeToCamel(s)
	t.Log(s)
}
