package mysql

import (
	"fmt"
	"github.com/pharosnet/dalc/cmd/dalc/internal/entry"
	"github.com/pharosnet/dalc/cmd/dalc/internal/parser/commons"
	"strings"
)

func ParseMySQLSchema(content string) (schema *entry.Schema, err error) {

	blocks := strings.Split(content, ";")

	schema = &entry.Schema{
		Name:   "",
		Tables: make([]*entry.Table, 0, 1),
	}
	changes := make([]*Change, 0, 1)

	for _, block := range blocks {

		lines := commons.NewLines(block)
		for lines.HasNext() {

			line := lines.NextLine()
			lineUpper := strings.ToUpper(line)

			// schema
			if strings.Index(lineUpper, "USE ") >= 0 {
				lines.Reset()
				schemaName, parseSchemaErr := parseMySQLSchema(lines)
				if parseSchemaErr != nil {
					err = parseSchemaErr
					return
				}
				schema.Name = schemaName
				continue
			}

			// table
			if strings.Index(lineUpper, "CREATE TABLE ") >= 0 {
				lines.Reset()
				table, parseTableErr := parseMySQLTable(lines)
				if parseTableErr != nil {
					err = parseTableErr
					return
				}
				if schema.Name == "" {
					if table.Schema == "" {
						err = fmt.Errorf("panic, no schema defined for table %s", table.Name)
						return
					}
					schema.Name = table.Schema
				} else {
					if table.Schema == "" {
						table.Schema = schema.Name
					}
					if schema.Name != table.Schema {
						err = fmt.Errorf("panic, can not define table %s in schema %s", table.FullName, schema.Name)
						return
					}
				}
				schema.Tables = append(schema.Tables, table)
				continue
			}

			// alter and DROP TABLE
			//if strings.Index(lineUpper, "ALTER TABLE ") >= 0 || strings.Index(lineUpper, "DROP TABLE ") >= 0 {
			//	lines.Reset()
			//	change, parseChangeErr := parseMySQLChange(lines)
			//	if parseChangeErr != nil {
			//		err = parseChangeErr
			//		return
			//	}
			//	if change.Schema == "" {
			//		change.Schema = schema.Name
			//	} else {
			//		if change.Schema != schema.Name {
			//			err = fmt.Errorf("panic, can not define change %s in schema %s", change.Content, schema.Name)
			//			return
			//		}
			//	}
			//	changes = append(changes, change)
			//	continue
			//}
		}

	}

	if schema.Name == "" {
		err = fmt.Errorf("no schema name got in schema defines")
		return
	}

	if len(schema.Tables) == 0 {
		err = fmt.Errorf("no tables got in schema defines")
		return
	}

	// todo
	if len(changes) > 0 {
		for _, change := range changes {
			for i := 0; i < len(schema.Tables); i++ {
				table := schema.Tables[i]
				if table.Name != change.Table {
					continue
				}
				switch change.Kind {
				case TableDropChangeKind:
					_, ok := change.Target.(*ChangDropTable)
					if ok {
						newTables := make([]*entry.Table, 0, 1)
						for _, table0 := range schema.Tables {
							if table0.Name != change.Table {
								newTables = append(newTables, table0)
							}
						}
						schema.Tables = newTables
					}
				case TableRenameChangeKind:
					rename, _ := change.Target.(*ChangeRenameTable)
					table.Name = rename.Name
					table.Schema = rename.Schema
				case ColumnAddChangeKind:
					add, _ := change.Target.(*ChangeAddColumn)
					newCols := make([]*entry.Column, 0, 1)
					newCol := &entry.Column{
						Name:         add.Name,
						Type:         add.Type,
						DefaultValue: add.DefaultValue,
					}
					if add.First {
						newCols = append(newCols, newCol)
						newCols = append(newCols, table.Columns...)
					} else {
						if add.After != "" {
							for _, column := range table.Columns {
								newCols = append(newCols, column)
								if column.Name == add.After {
									newCols = append(newCols, newCol)
								}
							}
						} else {
							newCols = append(newCols, table.Columns...)
							newCols = append(newCols, newCol)
						}
					}
					table.Columns = newCols
				case ColumnModifyChangeKind:
					mod, _ := change.Target.(*ChangeModifyColumn)
					for _, column := range table.Columns {
						if column.Name == mod.Source {
							column.Type = mod.Type
							column.DefaultValue = mod.DefaultValue
						}
					}
				case ColumnChangeChangeKind:
					ch, _ := change.Target.(*ChangeChangeColumn)
					for _, column := range table.Columns {
						if column.Name == ch.Source {
							column.Name = ch.Name
							column.Type = ch.Type
							column.DefaultValue = ch.DefaultValue
						}
					}
				case ColumnDropChangeKind:
					drop, _ := change.Target.(*ChangDropColumn)
					newCols := make([]*entry.Column, 0, 1)
					for _, column := range table.Columns {
						if column.Name != drop.Name {
							newCols = append(newCols, column)
						}
					}
					table.Columns = newCols
				default:
					err = fmt.Errorf("unknonw change %v", change.Kind)
					return
				}
			}
		}
	}

	return
}

func parseMySQLSchema(lines *commons.Lines) (schema string, err error) {
	for lines.HasNext() {
		line := lines.NextLine()
		lineUpper := strings.ToUpper(line)
		if strings.Index(lineUpper, "USE ") >= 0 {
			words := lines.CurrentLineWords()
			if len(words) > 1 {
				schema = commons.NormalizeName(words[1])
				break
			}
		}
	}
	if schema == "" {
		lines.Reset()
		line := lines.Remain()
		err = fmt.Errorf("read db schema failed at %s", line)
	}
	return
}

func parseMySQLTable(lines *commons.Lines) (table *entry.Table, err error) {
	structName := ""
	table = &entry.Table{
		FullName: "",
		Schema:   "",
		Name:     "",
		Columns:  make([]*entry.Column, 0, 1),
	}
	columnsBegin := false
	columnsEnd := false
	columnLineClean := false
	for lines.HasNext() {
		line := lines.NextLine()
		lineUpper := strings.ToUpper(line)

		// comments
		if structName == "" {
			if strings.Index(lineUpper, "--") >= 0 {
				words := lines.CurrentLineWords()
				for i, word := range words {
					if strings.ToUpper(word) == "NAME:" {
						structName = words[i+1]
						break
					}
				}
				if structName != "" {
					continue
				}
			}
		}

		// table name and schema
		if strings.Index(lineUpper, "CREATE TABLE ") >= 0 {
			words := lines.CurrentLineWords()
			fullTableName := ""
			for i, word := range words {
				word = strings.ToUpper(word)
				if word == "TABLE" {
					if words[i+1] == "IF" && words[i+2] == "NOT" && words[i+3] == "EXISTS" {
						fullTableName = words[i+4]
					} else {
						fullTableName = words[i+1]
					}
					break
				}
			}
			if fullTableName == "" {
				err = fmt.Errorf("read table name failed in %s", line)
				return
			}
			if structName == "" {
				err = fmt.Errorf("table %s has no defined struct name, please use -- name: {struct name}", fullTableName)
			}
			table.FullName = commons.NormalizeName(fullTableName)
			table.Schema, table.Name = commons.SplitFullName(fullTableName)
			lastWord := words[len(words)-1]
			if lastWord == "(" {
				columnsBegin = true
			}
			continue
		}

		// columns
		if !columnsBegin && !columnsEnd {
			if strings.Index(lineUpper, "(") >= 0 {
				words := commons.ReadWords([]byte(lineUpper))
				lastWord := words[len(words)-1]
				if lastWord != "(" {
					columnLineClean = true
					lines.Prev()
				}
				columnsBegin = true
			}
			continue
		}

		if columnsBegin && !columnsEnd {
			columnLine := strings.ReplaceAll(line, ",", " ")
			words := commons.ReadWords([]byte(columnLine))
			if columnLineClean {
				// remove ( in first column line
				words = words[1:]
				columnLineClean = false
			}
			if len(words) < 2 {
				err = fmt.Errorf("read table %s column failed in %s", table.FullName, line)
				return
			}
			columnsEnd = commons.WordsContains(words, "PRIMARY", "KEY", "UNIQUE")
			if columnsEnd {
				columnsBegin = false
				continue
			}
			columnName := commons.NormalizeName(words[0])
			columnType, columnTypeErr := entry.NewColumnType(words[1])
			if columnTypeErr != nil {
				err = fmt.Errorf("read table %s column failed, %s %s, %v in %s", table.FullName, columnName, words[1], columnTypeErr, line)
				return
			}
			defaultValue := ""
			defaultKeyIdx := commons.WordsIndex(words, "DEFAULT")
			if defaultKeyIdx > 0 {
				defaultValue = commons.NormalizeUpperValue(words[defaultKeyIdx+1])
			}
			var goType *entry.GoType
			refTypeKeyIdx := commons.WordsIndex(words, "REF:")
			if refTypeKeyIdx > 0 {
				refType := commons.NormalizeValue(words[refTypeKeyIdx+1])
				goType = entry.NewGoType(refType)
			}
			if goType == nil {
				goType0, goTypeErr := columnType.GoType()
				if goType0 == nil || goTypeErr != nil {
					err = fmt.Errorf("read table %s column failed, %s, %v in %s", table.FullName, columnName, goTypeErr, line)
					return
				}
				goType = goType0
			}

			column := &entry.Column{
				Name:         columnName,
				Type:         columnType,
				GoType:       goType,
				DefaultValue: defaultValue,
			}
			table.Columns = append(table.Columns, column)
		}

	}
	return
}

// not supported
func parseMySQLChange(lines *commons.Lines) (change *Change, err error) {
	content := lines.Remain()
	words := commons.ReadWords([]byte(strings.ToUpper(content)))
	if len(words) < 3 {
		err = fmt.Errorf("read change failed in \n%s", content)
		return
	}
	change = &Change{
		Kind:    "",
		Schema:  "",
		Table:   "",
		Target:  nil,
		Content: commons.WordsToLine(words),
	}
	change.Schema, change.Table = commons.SplitFullName(words[2])
	if words[0] == "DROP" {
		change.Kind = TableDropChangeKind
		change.Target = &ChangDropTable{}
	} else if words[0] == "ALTER" {
		// todo
		// DROP

	} else {
		err = fmt.Errorf("read change failed, unknonw change command %s, \n%s", words[0], content)
		return
	}

	return
}
