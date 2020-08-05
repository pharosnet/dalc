package commons_test

import (
	"github.com/pharosnet/dalc/v2/cmd/dalc/internal/parser/commons"
	"testing"
)

func TestWordsContainsAll(t *testing.T) {
	words := []string{"abc", "ddd", "---", "123"}
	has := commons.WordsContainsAll(words, "---", "ddd")
	t.Log(has)
	t.Log(commons.WordsContainsAll(words, "---", "xx"))
}
