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
			for _, table := range tables {
				if expr.Table.Schema == table.Schema && expr.Table.Table == table.Name {
					for _, column := range table.Columns {
						if column.Name == expr.ColumnName {
							expr.GoType = column.GoType
						}
					}
				}
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
			for _, table := range tables {
				if expr.Table.Schema == table.Schema && expr.Table.Table == table.Name {
					for _, column := range table.Columns {
						if column.Name == expr.ColumnName {
							expr.GoType = column.GoType
						}
					}
				}
			}
		}
	}
	queries = queries0
	return
}
