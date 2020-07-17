package config

import (
	"github.com/pharosnet/dalc/cmd/dalc/internal/logs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func NewConfig(dialect string, pkgName string, out string, enableJsonTags bool, schemaPath string, queryPath string) (conf *Config1, err error) {

	schemaContent, schemaContentErr := readFiles(schemaPath)
	if schemaContentErr != nil {
		err = schemaContentErr
		return
	}
	queryContent, queryContentErr := readFiles(queryPath)
	if queryContentErr != nil {
		err = queryContentErr
		return
	}

	conf = &Config1{
		Dialect:        dialect,
		Out:            out,
		PackageName:    pkgName,
		EnableJsonTags: enableJsonTags,
		SchemaFiles:    schemaContent,
		QueryFiles:     queryContent,
	}

	return
}

type Config1 struct {
	Pwd            string
	Dialect        string
	Out            string
	PackageName    string
	EnableJsonTags bool
	SchemaFiles    []*ContentedFile
	QueryFiles     []*ContentedFile
}

type ContentedFile struct {
	Name    string
	Content []byte
}

func readFiles(path string) (contentFiles []*ContentedFile, err error) {
	file, openErr := os.Open(path)
	if openErr != nil {
		err = openErr
		return
	}
	stat, statErr := file.Stat()
	if statErr != nil {
		err = statErr
		return
	}
	contentFiles = make([]*ContentedFile, 0, 1)
	if stat.IsDir() {
		files, readDirErr := ioutil.ReadDir(path)
		if readDirErr != nil {
			logs.Log().Println("read", path, "failed", readDirErr)
			err = readDirErr
			return
		}
		for _, file := range files {
			contentedFiles, readErr := readFiles(filepath.Join(path, file.Name()))
			if readErr != nil {
				err = readErr
				return
			}
			contentFiles = append(contentFiles, contentedFiles...)
		}
	} else {
		if strings.Contains(strings.ToLower(path), ".sql") {
			content, readErr := ioutil.ReadFile(path)
			if readErr != nil {
				err = readErr
				return
			}
			_, fileName := filepath.Split(path)
			contentFiles = append(contentFiles, &ContentedFile{
				Name:    fileName,
				Content: content,
			})
		}
	}
	return
}
