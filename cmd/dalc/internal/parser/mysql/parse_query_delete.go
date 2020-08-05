package mysql

import (
	"fmt"
	"github.com/pharosnet/dalc/v2/cmd/dalc/internal/entry"
	"github.com/vitessio/vitess/go/vt/sqlparser"
)

func parseQueryDelete(query *entry.Query, stmt *sqlparser.Delete) (err error) {

	query.TableList = append(query.TableList, &entry.QueryTable{
		Schema: stmt.Table.Qualifier.CompliantName(),
		Table:  stmt.Table.Name.CompliantName(),
		NameAs: "",
	})

	err = parseQueryWhere(query, stmt.Where)
	if err != nil {
		return
	}

	if len(query.CondExprList.ExprList) < 1 {
		err = fmt.Errorf("parse delete failed, found no condm in \n%s", query.Sql)
		return
	}

	return
}
