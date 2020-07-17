package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc"
	"time"
)

const (
	userRowGetByIdSQL = "SELECT `users`.`id`, `users`.`name`, `users`.`age`, `users`.`has_friends`, `users`.`join_time` FROM `ddd_test`.`users` WHERE `id` = ?"
	userRowInsertSQL  = "INSERT INTO `ddd_test`.`users` (`id`,`name`,`age`,`has_friends`,`join_time`) VALUES (?,?,?,?)"
	userRowUpdateSQL  = "UPDATE `ddd_test`.`users` SET `name` = ?,`age` = ?,`has_friends` = ?,`join_time` = ? WHERE `id` = ?"
	userRowDeleteSQL  = "DELETE FROM `ddd_test`.`users` WHERE `id` = ? "
)

type UserRow struct {
	Id         int64
	Name       string
	Age        int
	HasFriends bool
	JoinTime   time.Time
}

func (row *UserRow) scanSQLRow(rows *sql.Rows) (err error) {
	err = rows.Scan(
		&row.Id,
		&row.Name,
		&row.Age,
		&row.HasFriends,
		&row.JoinTime,
	)
	return
}

func (row *UserRow) conventToGetArgs() (args *dalc.Args) {
	args = dalc.NewArgs()
	args.Arg(row.Id)
	return
}

func (row *UserRow) Get(ctx dalc.PreparedContext) (err error) {
	err = dalc.Query(ctx, userRowGetByIdSQL, row.conventToGetArgs(), func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {
		if rowErr != nil {
			err = rowErr
			return
		}
		err = row.scanSQLRow(rows)
		return
	})
	return
}

func (row *UserRow) conventToInsertArgs() (args *dalc.Args) {
	args = dalc.NewArgs().Arg(row.Name).Arg(row.Age).Arg(row.HasFriends).Arg(row.JoinTime)
	return
}

func (row *UserRow) Insert(ctx dalc.PreparedContext) (err error) {
	insertId, execErr := dalc.Execute(ctx, userRowInsertSQL, row.conventToInsertArgs())
	if execErr != nil {
		err = execErr
		return
	}
	row.Id = insertId
	return
}

func (row *UserRow) conventToUpdateArgs() (args *dalc.Args) {
	args = dalc.NewArgs().Arg(row.Name).Arg(row.Age).Arg(row.HasFriends).Arg(row.JoinTime).Arg(row.Id)
	return
}

func (row *UserRow) Update(ctx dalc.PreparedContext) (err error) {
	_, execErr := dalc.Execute(ctx, userRowUpdateSQL, row.conventToUpdateArgs())
	if execErr != nil {
		err = execErr
		return
	}
	return
}

func (row *UserRow) conventToDeleteArgs() (args *dalc.Args) {
	args = dalc.NewArgs().Arg(row.Id)
	return
}

func (row *UserRow) Delete(ctx dalc.PreparedContext) (err error) {
	_, execErr := dalc.Execute(ctx, userRowDeleteSQL, row.conventToDeleteArgs())
	if execErr != nil {
		err = execErr
		return
	}
	return
}
