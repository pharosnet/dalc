package generates

import "github.com/pharosnet/dalc/cmd/dalc/internal/entry"

func Generate(out string, tables []*entry.Table, queries []*entry.Query) (err error) {

	fs := make([]*GenerateFile,0,1)
	fs0, tableErr := generateTable(tables)
	if tableErr != nil {
		err = tableErr
		return
	}
	fs = append(fs, fs0...)
	fs1, queryErr := generateQuery(queries)
	if queryErr != nil {
		err = queryErr
		return
	}
	fs = append(fs, fs1...)

	err = writeFiles(out, fs)

	return
}
