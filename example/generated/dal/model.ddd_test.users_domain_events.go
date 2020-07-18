package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc"
)

const (
	usersDomainEventsRowGetByPkSQL = "SELECT `ID`, `AGGREGATE_NAME`, `AGGREGATE_ID`, `EVENT_NAME`, `EVENT_ID`, `EVENT_TIME`, `EVENT_DATA`FROM `DDD_TEST`.`USERS_DOMAIN_EVENTS` WHERE `ID` = ?"
	usersDomainEventsRowInsertSQL  = "INSERT INTO `DDD_TEST`.`USERS_DOMAIN_EVENTS` ( `AGGREGATE_NAME`, `AGGREGATE_ID`, `EVENT_NAME`, `EVENT_ID`, `EVENT_TIME`, `EVENT_DATA`) VALUES ( ?, ?, ?, ?, ?, ?)"
	usersDomainEventsRowUpdateSQL  = "UPDATE `DDD_TEST`.`USERS_DOMAIN_EVENTS` SET `AGGREGATE_NAME` = ?, `AGGREGATE_ID` = ?, `EVENT_NAME` = ?, `EVENT_ID` = ?, `EVENT_TIME` = ?, `EVENT_DATA` = ? WHERE `ID` = ?"
	usersDomainEventsRowDeleteSQL  = "DELETE FROM `DDD_TEST`.`USERS_DOMAIN_EVENTS` WHERE `ID` = ?"
)

type UsersDomainEventsRow struct {
	Id          int64          `json:"id"`
	AggName     string         `json:"agg_name"`
	AggregateId string         `json:"aggregate_id"`
	EventName   string         `json:"event_name"`
	EventId     string         `json:"event_id"`
	EventTime   dalc.MySQLTime `json:"event_time"`
	EventData   dalc.NullBytes `json:"event_data"`
}

func (row *UsersDomainEventsRow) scanSQLRow(rows *sql.Rows) (err error) {
	err = rows.Scan(
		&row.Id,
		&row.AggName,
		&row.AggregateId,
		&row.EventName,
		&row.EventId,
		&row.EventTime,
		&row.EventData,
	)
	return
}

func (row *UsersDomainEventsRow) conventToGetArgs() (args *dalc.Args) {

	args = dalc.NewArgs()
	args.Arg(row.Id)

	return
}

func (row *UsersDomainEventsRow) Get(ctx dalc.PreparedContext) (err error) {
	err = dalc.Query(ctx, usersDomainEventsRowGetByPkSQL, row.conventToGetArgs(), func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {
		if rowErr != nil {
			err = rowErr
			return
		}
		err = row.scanSQLRow(rows)
		return
	})
	return
}

func (row *UsersDomainEventsRow) conventToInsertArgs() (args *dalc.Args) {

	args = dalc.NewArgs()

	args.Arg(row.AggName)
	args.Arg(row.AggregateId)
	args.Arg(row.EventName)
	args.Arg(row.EventId)
	args.Arg(row.EventTime)
	args.Arg(row.EventData)

	return
}

func (row *UsersDomainEventsRow) Insert(ctx dalc.PreparedContext) (err error) {

	insertId, execErr := dalc.ExecuteReturnInsertId(ctx, usersDomainEventsRowInsertSQL, row.conventToInsertArgs())
	if execErr != nil {
		err = execErr
		return
	}
	row.Id = insertId

	return
}

func (row *UsersDomainEventsRow) conventToUpdateArgs() (args *dalc.Args) {

	args = dalc.NewArgs()

	args.Arg(row.AggName)
	args.Arg(row.AggregateId)
	args.Arg(row.EventName)
	args.Arg(row.EventId)
	args.Arg(row.EventTime)
	args.Arg(row.EventData)

	args.Arg(row.Id)

	return
}

func (row *UsersDomainEventsRow) Update(ctx dalc.PreparedContext) (err error) {
	_, execErr := dalc.Execute(ctx, usersDomainEventsRowUpdateSQL, row.conventToUpdateArgs())
	if execErr != nil {
		err = execErr
		return
	}
	return
}

func (row *UsersDomainEventsRow) conventToDeleteArgs() (args *dalc.Args) {

	args = dalc.NewArgs()
	args.Arg(row.Id)

	return
}

func (row *UsersDomainEventsRow) Delete(ctx dalc.PreparedContext) (err error) {
	_, execErr := dalc.Execute(ctx, usersDomainEventsRowDeleteSQL, row.conventToDeleteArgs())
	if execErr != nil {
		err = execErr
		return
	}
	return
}