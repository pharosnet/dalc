package entry

import "fmt"

func QueryFill(tables []*Table, queries0 []*Query) (queries []*Query, err error) {
	for _, query := range queries0 {

		for _, queryTable := range query.TableList {
			for _, table := range tables {
				if queryTable.Schema == table.Schema && queryTable.Table == table.Name {
					queryTable.Ref = table
				}
			}
		}

		for _, expr := range query.SelectExprList.ExprList {
			expr.BuildName()
			if expr.Name == "" {
				err = fmt.Errorf("fill %s failed, one select expr has go type, but no name", query.Name)
				return
			}
			if expr.GoType != nil {
				continue
			}
			got := false
			for _, table := range tables {
				if got {
					break
				}
				if expr.Table.Schema == table.Schema && expr.Table.Table == table.Name {
					for _, column := range table.Columns {
						if column.Name == expr.ColumnName {
							expr.GoType = column.GoType
							got = true
							break
						}
					}
				}
			}
			if !got {
				err = fmt.Errorf("fill %s failed, select expr %s does not find type", query.Name, expr.Name)
			}
		}

		for _, expr := range query.CondExprList.ExprList {
			expr.BuildName()
			if expr.Name == "" {
				err = fmt.Errorf("fill %s failed, one select expr has go type, but no name", query.Name)
				return
			}
			if expr.GoType != nil {
				continue
			}
			got := false
			for _, table := range tables {
				if got {
					break
				}
				if expr.Table.Schema == table.Schema && expr.Table.Table == table.Name {
					for _, column := range table.Columns {
						if column.Name == expr.ColumnName {
							expr.GoType = column.GoType
							got = true
							break
						}
					}
				}
			}
			if !got {
				err = fmt.Errorf("fill %s failed, where cond expr %s does not find type", query.Name, expr.Name)
			}
		}
	}
	queries = queries0
	return
}
