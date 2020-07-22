package commons_test

import (
	"github.com/pharosnet/dalc/cmd/dalc/v2/internal/parser/commons"
	"testing"
)

func TestWordsContainsAll(t *testing.T) {
	words := []string{"abc", "ddd", "---", "123"}
	has := commons.WordsContainsAll(words, "---", "ddd")
	t.Log(has)
	t.Log(commons.WordsContainsAll(words, "---", "xx"))
}
