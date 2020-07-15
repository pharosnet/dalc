package mysql

import (
	"fmt"
	"github.com/pharosnet/dalc/cmd/dalc/internal/entry"
	"github.com/pharosnet/dalc/cmd/dalc/internal/parser/commons"
	"github.com/vitessio/vitess/go/vt/sqlparser"
	"reflect"
	"strings"
)

func parseSelectExprs(query *entry.Query, stmt sqlparser.SelectExprs) (err error) {

	err = stmt.WalkSubtree(func(node sqlparser.SQLNode) (kontinue bool, err error) {

		switch node.(type) {
		case *sqlparser.StarExpr:
			err = fmt.Errorf("parse query exprs failed, star(*) expr is not support, \n%s", query.Sql)
		case *sqlparser.NonStarExpr:
			err = parseSelectNonStarExpr(query, node.(*sqlparser.NonStarExpr))
		default:
			err = fmt.Errorf("parse query exprs failed, %v is not support, \n%s", reflect.TypeOf(node), query.Sql)
		}

		return
	})

	return
}

func parseSelectNonStarExpr(query *entry.Query, stmt *sqlparser.NonStarExpr) (err error) {

	queryExpr := &entry.QueryExpr{
		Table:               nil,
		ColumnQualifierName: "",
		ColumnName:          "",
		FuncName:            "",
		Name:                "",
		GoType:              nil,
	}
	queryExprGot := false

	err = stmt.WalkSubtree(func(node sqlparser.SQLNode) (kontinue bool, err error) {

		switch node.(type) {
		case *sqlparser.ColName:
			if queryExpr.ColumnName == "" {
				x := node.(*sqlparser.ColName)
				queryExpr.ColumnName = x.Name.CompliantName()
				if x.Qualifier != nil {
					queryExpr.ColumnQualifierName = x.Qualifier.Name.CompliantName()
				}
				queryExprGot = true
			}

		case sqlparser.ColIdent:
			if queryExpr.Name == "" {
				ident := node.(sqlparser.ColIdent)
				identName := ident.CompliantName()
				if identName != "" {
					queryExpr.Name = commons.SnakeToCamel(identName)
				}
			}
		case sqlparser.Expr:

		case *sqlparser.FuncExpr:
			x := node.(*sqlparser.FuncExpr)
			err = parseSelectNonStarFunExpr(queryExpr, x)
		case *sqlparser.Subquery:
			x := node.(*sqlparser.Subquery)
			subQuery := entry.NewQuery()
			err = parseQuerySelect(subQuery, x.Select)
			if err != nil {
				return
			}
			subQuery.Fill()
			if len(subQuery.SelectExprList.ExprList) != 1 {
				err = fmt.Errorf("parse non star expr failed, sub query must has one expr, but found %d", len(subQuery.SelectExprList.ExprList))
				return
			}
			subExpr := subQuery.SelectExprList.ExprList[0]
			queryExpr.Table = subExpr.Table
			queryExpr.ColumnName = subExpr.ColumnName
			queryExpr.FuncName = subExpr.FuncName
			queryExpr.GoType = subExpr.GoType
			entry.QueryMergeCond(query, subQuery)

		case *sqlparser.ParenExpr: // ( ... )

		case *sqlparser.ComparisonExpr: // left ? right

		case *sqlparser.AndExpr: // left and right

		case *sqlparser.OrExpr: // left or right

		case *sqlparser.NotExpr: // not ?


		default:
			err = fmt.Errorf("parse non star expr failed, %v is not support", reflect.TypeOf(node))
		}

		queryExpr.BuildName()
		query.SelectExprList.ExprList = append(query.SelectExprList.ExprList, queryExpr)
		return
	})

	if err != nil {
		err = fmt.Errorf("parse non star expr failed, %v, \n%s", err, query.Sql)
	}

	return
}

func parseSelectNonStarFunExpr(expr *entry.QueryExpr, stmt *sqlparser.FuncExpr) (err error) {
	funName := strings.ToLower(stmt.Name.CompliantName())
	if funName == "count" {
		expr.GoType = entry.NewGoType("int")
	}
	expr.FuncName = funName
	err = stmt.WalkSubtree(func(node sqlparser.SQLNode) (kontinue bool, err error) {
		switch node.(type) {
		case *sqlparser.NonStarExpr:
			nonStar := node.(*sqlparser.NonStarExpr)
			_ = nonStar.WalkSubtree(func(node sqlparser.SQLNode) (kontinue bool, err error) {
				switch node.(type) {
				case *sqlparser.ColName:
					x := node.(*sqlparser.ColName)
					expr.ColumnName = x.Name.CompliantName()
					if x.Qualifier != nil {
						expr.ColumnQualifierName = x.Qualifier.Name.CompliantName()
					}
				case sqlparser.ColIdent:

				default:
					err = fmt.Errorf("parse non star func expr failed %v is not support", reflect.TypeOf(node))
				}

				return
			})
		case sqlparser.ColIdent:
			ident := node.(sqlparser.ColIdent)
			identName := ident.CompliantName()
			if identName != "" {
				expr.Name = commons.SnakeToCamel(identName)
			}
		default:
			err = fmt.Errorf("parse non star func expr failed, %v, %v is not support", funName, reflect.TypeOf(node))
		}
		return
	})

	return
}

func parseSelectSubQueryExpr(expr *entry.QueryExpr, stmt *sqlparser.Subquery) (err error) {

	err = stmt.WalkSubtree(func(node sqlparser.SQLNode) (kontinue bool, err error) {
		switch node.(type) {
		case sqlparser.Comments:

		case sqlparser.SelectExprs:
			x := node.(sqlparser.SelectExprs)
			err = x.WalkSubtree(func(node sqlparser.SQLNode) (kontinue bool, err error) {

				return
			})
			err = parseSelectExprs(query, node.(sqlparser.SelectExprs))
		case sqlparser.TableExprs:

		case *sqlparser.Where:

		case *sqlparser.Limit:

		case sqlparser.GroupBy:

		case sqlparser.OrderBy:

		default:
			err = fmt.Errorf("parse query select failed, %s is not support, in \n%s", reflect.TypeOf(node), query.Sql)
		}

		return
	})

	return
}
