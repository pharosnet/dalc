package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc"
)

// ************* users_domain_events_list *************
const usersDomainEventsListSQL = "SELECT `ee`.`id` as `xxxx`, `ee`.`aggregate_name`, `users_domain_events`.`aggregate_id`, `ee`.`event_name`, `ee`.`event_id` FROM `ddd_test`.`users_domain_events` as `ee` WHERE `ee`.`aggregate_id` = ? AND `ee`.`aggregate_name` = 'DD' AND `ee`.`event_id` IN ('#xxxx#') and `ee`.`event_name` between ? and ? ORDER BY `ee`.`id` DESC LIMIT ? OFFSET ? "

type UsersDomainEventsListRequest struct {
	AggregateId string
	EventIdList []string
	EventName   string
	Offset      int
	Limit       int
}

type UsersDomainEventsListResult struct {
	XXXX          int64  `json:"xxxx"`
	AggregateName string `json:"aggregate_name"`
	AggregateId   string `json:"aggregate_id"`
	EventName     string `json:"event_name"`
	EventId       string `json:"event_id"`
}

type UsersDomainEventsListResultIterator func(ctx context.Context, result *UsersDomainEventsListResult) (err error)

func UsersDomainEventsList(ctx dalc.PreparedContext, request *UsersDomainEventsListRequest, iterator UsersDomainEventsListResultIterator) (err error) {

	querySQL := usersDomainEventsListSQL
	args := dalc.NewArgs()
	args.Arg(request.AggregateId)

	querySQL = dalc.ReplaceSQL(querySQL, "#XXXX#", dalc.NewTupleArgs(request.EventIdList))

	args.Arg(request.EventName)
	args.Arg(request.Offset)
	args.Arg(request.Limit)

	err = dalc.Query(ctx, querySQL, args, func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {

		if rowErr != nil {
			err = rowErr
			return
		}

		result := &UsersDomainEventsListResult{}
		scanErr := rows.Scan(
			&result.XXXX,
			&result.AggregateName,
			&result.AggregateId,
			&result.EventName,
			&result.EventId,
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
