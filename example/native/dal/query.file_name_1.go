package dal

import (
	"context"
	"database/sql"
	"github.com/pharosnet/dalc/v2"
	"time"
)

// ************* UserListByJoinTime *************
const userListByJoinTimeSQL = "SELECT `users`.`id`, `users`.`name`, `users`.`age`, `users`.`has_friends`, `users`.`join_time` FROM `ddd_test`.`users` WHERE `id` = ?"

type UserListByJoinTimeRequest struct {
	Id []int64
}

type UserListByJoinTimeResult struct {
	Id         int64
	Name       string
	Age        int
	HasFriends bool
	JoinTime   time.Time
}

type UserListByJoinTimeResultIterator func(ctx context.Context, result *UserListByJoinTimeResult) (err error)

func UserListByJoinTime(ctx dalc.PreparedContext, request *UserListByJoinTimeRequest, iterator UserListByJoinTimeResultIterator) (err error) {

	args := dalc.NewArgs()
	args.Arg(request.Id)

	err = dalc.Query(ctx, userListByJoinTimeSQL, args, func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {

		if rowErr != nil {
			err = rowErr
			return
		}
		result := &UserListByJoinTimeResult{}

		scanErr := rows.Scan(
			&result.Id,
			&result.Name,
			&result.Age,
			&result.HasFriends,
			&result.JoinTime,
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

// ************* UserUpdateName *************
const userUpdateNameSQL = "UPDATE `ddd_test`.`users` SET `join_time` = ? WHERE `id` = ?"

type UserUpdateNameRequest struct {
	Id       int64
	JoinTime time.Time
}

func UserUpdateName(ctx dalc.PreparedContext, request *UserUpdateNameRequest) (affected int64, err error) {

	args := dalc.NewArgs()
	args.Arg(request.JoinTime)
	args.Arg(request.Id)

	affected, err = dalc.Execute(ctx, userUpdateNameSQL, args)

	return
}
