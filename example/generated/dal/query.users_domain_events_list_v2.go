package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc"
)

// ************* users_domain_events_list_v2 *************
const usersDomainEventsListV2SQL = "SELECT e.id, s.aggregate_id as `snapshot_aggregate_id`, e.aggregate_id as `event_aggregate_id` FROM ddd_test.users_domain_events as e left join ddd_test.users_domain_snapshot as s on s.aggregate_id = e.aggregate_id left join ddd_test.users_domain_snapshot as s1 on s1.aggregate_id = e.aggregate_id where e.id = ? and s.aggregate_id = ? "

type UsersDomainEventsListV2Request struct {
	Id          int64
	AggregateId sql.NullString
}

type UsersDomainEventsListV2Result struct {
	Id                    int64          `json:"id"`
	SNAPSHOT_AGGREGATE_ID sql.NullString `json:"snapshot__aggregate__id"`
	EVENT_AGGREGATE_ID    string         `json:"event__aggregate__id"`
}

type UsersDomainEventsListV2ResultIterator func(ctx context.Context, result *UsersDomainEventsListV2Result) (err error)

func UsersDomainEventsListV2(ctx dalc.PreparedContext, request *UsersDomainEventsListV2Request, iterator UsersDomainEventsListV2ResultIterator) (err error) {

	querySQL := usersDomainEventsListV2SQL
	args := dalc.NewArgs()
	args.Arg(request.Id)
	args.Arg(request.AggregateId)

	err = dalc.Query(ctx, querySQL, args, func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {

		if rowErr != nil {
			err = rowErr
			return
		}

		result := &UsersDomainEventsListV2Result{}
		scanErr := rows.Scan(
			&result.Id,
			&result.SNAPSHOT_AGGREGATE_ID,
			&result.EVENT_AGGREGATE_ID,
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
