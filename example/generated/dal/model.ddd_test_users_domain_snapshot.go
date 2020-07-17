package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc"
)

const (
	usersDomainSnapshotRowGetByPkSQL = "SELECT `ID`, `AGGREGATE_NAME`, `AGGREGATE_ID`, `LAST_EVENT_ID`, `SNAPSHOT_DATA`FROM `DDD_TEST`.`USERS_DOMAIN_SNAPSHOT` WHERE `ID` = ?"
	usersDomainSnapshotRowInsertSQL  = "INSERT INTO `DDD_TEST`.`USERS_DOMAIN_SNAPSHOT` ( `AGGREGATE_NAME`, `AGGREGATE_ID`, `LAST_EVENT_ID`, `SNAPSHOT_DATA`) VALUES ( ?, ?, ?, ?)"
	usersDomainSnapshotRowUpdateSQL  = "UPDATE `DDD_TEST`.`USERS_DOMAIN_SNAPSHOT` SET `AGGREGATE_NAME` = ?, `AGGREGATE_ID` = ?, `LAST_EVENT_ID` = ?, `SNAPSHOT_DATA` = ? WHERE `ID` = ?"
	usersDomainSnapshotRowDeleteSQL  = "DELETE FROM `DDD_TEST`.`USERS_DOMAIN_SNAPSHOT` WHERE `ID` = ?"
)

type UsersDomainSnapshotRow struct {
	Id            int64          `json:"id"`
	AggregateName sql.NullString `json:"aggregate_name"`
	AggregateId   sql.NullString `json:"aggregate_id"`
	LastEventId   sql.NullString `json:"last_event_id"`
	SnapshotData  dalc.NullBytes `json:"snapshot_data"`
}

func (row *UsersDomainSnapshotRow) scanSQLRow(rows *sql.Rows) (err error) {
	err = rows.Scan(
		&row.Id,
		&row.AggregateName,
		&row.AggregateId,
		&row.LastEventId,
		&row.SnapshotData,
	)
	return
}

func (row *UsersDomainSnapshotRow) conventToGetArgs() (args *dalc.Args) {

	args = dalc.NewArgs()
	args.Arg(row.Id)

	return
}

func (row *UsersDomainSnapshotRow) Get(ctx dalc.PreparedContext) (err error) {
	err = dalc.Query(ctx, usersDomainSnapshotRowGetByPkSQL, row.conventToGetArgs(), func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {
		if rowErr != nil {
			err = rowErr
			return
		}
		err = row.scanSQLRow(rows)
		return
	})
	return
}

func (row *UsersDomainSnapshotRow) conventToInsertArgs() (args *dalc.Args) {

	args = dalc.NewArgs()

	args.Arg(row.AggregateName)
	args.Arg(row.AggregateId)
	args.Arg(row.LastEventId)
	args.Arg(row.SnapshotData)

	return
}

func (row *UsersDomainSnapshotRow) Insert(ctx dalc.PreparedContext) (err error) {

	insertId, execErr := dalc.ExecuteReturnInsertId(ctx, usersDomainSnapshotRowInsertSQL, row.conventToInsertArgs())
	if execErr != nil {
		err = execErr
		return
	}
	row.Id = insertId

	return
}

func (row *UsersDomainSnapshotRow) conventToUpdateArgs() (args *dalc.Args) {

	args = dalc.NewArgs()

	args.Arg(row.AggregateName)
	args.Arg(row.AggregateId)
	args.Arg(row.LastEventId)
	args.Arg(row.SnapshotData)

	args.Arg(row.Id)

	return
}

func (row *UsersDomainSnapshotRow) Update(ctx dalc.PreparedContext) (err error) {
	_, execErr := dalc.Execute(ctx, usersDomainSnapshotRowUpdateSQL, row.conventToUpdateArgs())
	if execErr != nil {
		err = execErr
		return
	}
	return
}

func (row *UsersDomainSnapshotRow) conventToDeleteArgs() (args *dalc.Args) {

	args = dalc.NewArgs()
	args.Arg(row.Id)

	return
}

func (row *UsersDomainSnapshotRow) Delete(ctx dalc.PreparedContext) (err error) {
	_, execErr := dalc.Execute(ctx, usersDomainSnapshotRowDeleteSQL, row.conventToDeleteArgs())
	if execErr != nil {
		err = execErr
		return
	}
	return
}
