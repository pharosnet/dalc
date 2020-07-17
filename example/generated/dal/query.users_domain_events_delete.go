package dal

import (
	"github.com/pharosnet/dalc"
)

// ************* users_domain_events_delete *************
const usersDomainEventsDeleteSQL = "DELETE FROM `ddd_test`.`users_domain_events` WHERE `id` = ? "

type UsersDomainEventsDeleteRequest struct {
	Id int64
}

func UsersDomainEventsDelete(ctx dalc.PreparedContext, request *UsersDomainEventsDeleteRequest) (affected int64, err error) {

	querySQL := usersDomainEventsDeleteSQL
	args := dalc.NewArgs()
	args.Arg(request.Id)

	affected, err = dalc.Execute(ctx, querySQL, args)

	return
}
