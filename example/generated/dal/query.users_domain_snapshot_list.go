package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc"
)

// ************* users_domain_snapshot_list *************
const usersDomainSnapshotListSQL = "SELECT `users_domain_snapshot`.`id`, `users_domain_snapshot`.`aggregate_name`, /* dd */ `users_domain_snapshot`.`aggregate_id`, `users_domain_snapshot`.`last_event_id`, `users_domain_snapshot`.`snapshot_data`, (`users_domain_snapshot`.`id` > 1)                                                          as `over`, (select count(`id`) from `ddd_test`.`users_domain_events` where `users_domain_events`.`aggregate_id` = `users_domain_snapshot`.`aggregate_id`)       as `count`, (select sum(`id`) from `ddd_test`.`users_domain_events` where `users_domain_events`.`aggregate_id` = `users_domain_snapshot`.`aggregate_id`)       as `sum`, exists(select `id` from `ddd_test`.`users_domain_events` where `users_domain_events`.`aggregate_id` = `users_domain_snapshot`.`aggregate_id`) as `x` FROM `ddd_test`.`users_domain_snapshot` where `id` = ? "

type UsersDomainSnapshotListRequest struct {
	Id int64
}

type UsersDomainSnapshotListResult struct {
	Id            int64          `json:"id"`
	AggregateName sql.NullString `json:"aggregate_name"`
	AggregateId   sql.NullString `json:"aggregate_id"`
	LastEventId   sql.NullString `json:"last_event_id"`
	SnapshotData  dalc.NullBytes `json:"snapshot_data"`
	OVER          bool           `json:"over"`
	COUNT         int            `json:"count"`
	SUM           int64          `json:"sum"`
	X             bool           `json:"x"`
}

type UsersDomainSnapshotListResultIterator func(ctx context.Context, result *UsersDomainSnapshotListResult) (err error)

func UsersDomainSnapshotList(ctx dalc.PreparedContext, request *UsersDomainSnapshotListRequest, iterator UsersDomainSnapshotListResultIterator) (err error) {

	args := dalc.NewArgs()
	args.Arg(request.Id)

	err = dalc.Query(ctx, usersDomainSnapshotListSQL, args, func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {

		if rowErr != nil {
			err = rowErr
			return
		}

		result := &UsersDomainSnapshotListResult{}
		scanErr := rows.Scan(
			&result.Id,
			&result.AggregateName,
			&result.AggregateId,
			&result.LastEventId,
			&result.SnapshotData,
			&result.OVER,
			&result.COUNT,
			&result.SUM,
			&result.X,
		)

		if scanErr != nil {
			err = scanErr
			return
		}

		err = iterator(ctx, result)

		return
	})

	return
}
