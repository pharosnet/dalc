// DO NOT EDIT THIS FILE, IT IS GENERATED BY DALC
package dal

import (
	"github.com/pharosnet/dalc"
)

// ************* users_domain_events_insert *************
const usersDomainEventsInsertSQL = "INSERT INTO `ddd_test`.`users_domain_events` (`id`, `aggregate_name`, `aggregate_id`, `event_name`, `event_id`, `event_time`, `event_data`) VALUES (?,?,?,?,?,now(),'') "

type UsersDomainEventsInsertRequest struct {
	Id            int64
	AggregateName string
	AggregateId   string
	EventName     string
	EventId       string
}

func UsersDomainEventsInsert(ctx dalc.PreparedContext, request *UsersDomainEventsInsertRequest) (affected int64, err error) {

	querySQL := usersDomainEventsInsertSQL
	args := dalc.NewArgs()
	args.Arg(request.Id)
	args.Arg(request.AggregateName)
	args.Arg(request.AggregateId)
	args.Arg(request.EventName)
	args.Arg(request.EventId)

	affected, err = dalc.Execute(ctx, querySQL, args)

	return
}
