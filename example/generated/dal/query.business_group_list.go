// DO NOT EDIT THIS FILE, IT IS GENERATED BY DALC
package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc/v2"
)

// ************* business_group_list *************
const businessGroupListSQL = "SELECT `business_group`.`id`, `business_group`.`create_by`, `business_group`.`create_at`, `business_group`.`modify_by`, `business_group`.`modify_at`, `business_group`.`delete_by`, `business_group`.`delete_at`, `business_group`.`version`, `business_group`.`code`, `business_group`.`name`, `business_group`.`description` FROM `applications`.`business_group` ORDER BY `code` LIMIT ? OFFSET ? "

type BusinessGroupListRequest struct {
	Offset int
	Limit  int
}

type BusinessGroupListResultIterator func(ctx context.Context, result *BusinessGroupRow) (err error)

func BusinessGroupList(ctx dalc.PreparedContext, request *BusinessGroupListRequest, iterator BusinessGroupListResultIterator) (err error) {

	querySQL := businessGroupListSQL
	args := dalc.NewArgs()
	args.Arg(request.Offset)
	args.Arg(request.Limit)

	err = dalc.Query(ctx, querySQL, args, func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {

		if rowErr != nil {
			err = rowErr
			return
		}

		result := &BusinessGroupRow{}
		scanErr := result.scanSQLRow(rows)

		if scanErr != nil {
			err = scanErr
			return
		}

		err = iterator(ctx, result)

		return
	})

	return
}
