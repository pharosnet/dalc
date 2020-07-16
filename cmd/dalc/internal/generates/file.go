package generates

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
)

type GenerateFile struct {
	Name    string
	Content []byte
}

func writeFiles(out string, fs []*GenerateFile) (err error) {

	for _, f := range fs {

		formatted, fErr := format.Source(f.Content)
		if fErr != nil {
			err = fmt.Errorf("go fmt %s failed, %v, code:\n%s", f.Name, fErr, string(f.Content))
			return
		}
		f.Content = formatted

	}

	for _, f := range fs {
		
		fileName := filepath.Join(out, f.Name)

		wErr := ioutil.WriteFile(fileName, f.Content, os.ModePerm)
		if wErr != nil {
			err = fmt.Errorf("write file %s filed, %v", f.Name, wErr)
			return
		}

	}

	return
}
