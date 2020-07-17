package dal

import (
	"github.com/pharosnet/dalc"
)

// ************* users_domain_events_update *************
const usersDomainEventsUpdateSQL = "UPDATE `ddd_test`.`users_domain_events` SET `aggregate_name` = ?, `aggregate_id` = ?, `event_name` = '1', `event_id` = ?, `event_time` = now(), `event_data` = null WHERE `id` = ? "

type UsersDomainEventsUpdateRequest struct {
	AggregateName string
	AggregateId   string
	EventId       string
	Id            int64
}

func UsersDomainEventsUpdate(ctx dalc.PreparedContext, request *UsersDomainEventsUpdateRequest) (affected int64, err error) {

	querySQL := usersDomainEventsUpdateSQL
	args := dalc.NewArgs()
	args.Arg(request.AggregateName)
	args.Arg(request.AggregateId)
	args.Arg(request.EventId)
	args.Arg(request.Id)

	affected, err = dalc.Execute(ctx, querySQL, args)

	return
}
