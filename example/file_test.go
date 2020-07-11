package example_test

import (
	"io/ioutil"
	"testing"
)

func Test_ReadDir(t *testing.T) {

	path := `D:\images\GS2\122018_1\24a01-wa.png`
	files, readErr := ioutil.ReadDir(path)
	if readErr != nil {
		t.Error(readErr)
		return
	}

	for _, file := range files {
		t.Log(file.Name())
	}
}
