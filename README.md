# DALC
Database access layer for go.

## Feature

- Simple.
- No-reflect cost.
- Using callback function to decrease range times.
- Expandability.
- It is more convenient to use with dalc command.

## Usage

`go get -u github.com/pharosnet/dalc/v2`


```go
// ************* UserListByJoinTime *************

const userListByJoinTimeSQL = "SELECT `users`.`id`, `users`.`name`, `users`.`age`, `users`.`has_friends`, `users`.`join_time` FROM `ddd_test`.`users` WHERE `id` = ?"

type UserListByJoinTimeRequest struct {
	Id int64
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

```
PreparedStatement
```go
// select
var db *sql.DB
ctx := dalc.WithPreparedStatement(context.TODO(), db)
err := UserListByJoinTime(ctx, &UserListByJoinTimeRequest{Id: 1}, func(ctx context.Context, result *UserListByJoinTimeResult) (err error) {
    // handle result
    return
})
```
```go
// exec with tx
tx, _ := db.Begin()
ctx := dalc.WithPreparedStatement(context.TODO(), tx)
_, err := UserUpdateName(ctx, &UserUpdateNameRequest{Id: 1, JoinTime: time.Now()})

tx.Commit()
```
## Code Generates
### Install dalc 
`go get -u github.com/pharosnet/dalc/v2/cmd/dalc/v2`
### Write sql schema files
Write sql schema files in some folder, such as schema/, and foo.sql in this folder.
```sql
USE `ddd_test`;

-- name: users_domain_events
CREATE TABLE `users_domain_events`
(
    `id`             bigint       NOT NULL AUTO_INCREMENT,
    `aggregate_name` varchar(255) NOT NULL, -- name: AggName ref: github.com/foo/bar.SQLString
    `aggregate_id`   varchar(255) NOT NULL,
    `event_name`     varchar(255) NOT NULL,
    `event_id`       varchar(63)  NOT NULL,
    `event_time`     datetime(6) DEFAULT NULL,
    `event_data`     text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_ix_event_id` (`event_id`),
    KEY `users_ix_aggname` (`aggregate_name`),
    KEY `users_ix_aggid` (`aggregate_id`),
    KEY `users_ix_event_name` (`event_name`),
    KEY `users_ix_event_time` (`event_time` DESC)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
```
* first comment line is to define one name to replace table name, use `name: ` mark. 
* In columns, can use `name: ` define a new name and use `ref: ` define the go type.
### Write sql query files
Write sql query files in some folder, such as query/, and bar.sql in this folder.
```sql
-- name: users_domain_events_list
SELECT `ee`.`id` as `xxxx`,
       `ee`.`aggregate_name`,
       `users_domain_events`.`aggregate_id`,
       `ee`.`event_name`,
       `ee`.`event_id`
FROM `ddd_test`.`users_domain_events` as `ee`
WHERE `ee`.`aggregate_id` = ?
  AND `ee`.`aggregate_name` = 'DD'
  AND `ee`.`event_id` IN ('#xxxx#')
  and `ee`.`event_name` between ? and ?
ORDER BY `ee`.`id` DESC LIMIT ? OFFSET ?;

-- name: users_domain_snapshot_list
SELECT `users_domain_snapshot`.`id`,
       `users_domain_snapshot`.`aggregate_name`, /* dd */
       `users_domain_snapshot`.`aggregate_id`,
       `users_domain_snapshot`.`last_event_id`,
       `users_domain_snapshot`.`snapshot_data`,
       (`users_domain_snapshot`.`id` > 1)                                                          as `over`,
       (select count(`id`)
        from `ddd_test`.`users_domain_events`
        where `users_domain_events`.`aggregate_id` = `users_domain_snapshot`.`aggregate_id`)       as `count`,
       (select sum(`id`)
        from `ddd_test`.`users_domain_events`
        where `users_domain_events`.`aggregate_id` = `users_domain_snapshot`.`aggregate_id`)       as `sum`,
       exists(select `id`
              from `ddd_test`.`users_domain_events`
              where `users_domain_events`.`aggregate_id` = `users_domain_snapshot`.`aggregate_id`) as `x`
FROM `ddd_test`.`users_domain_snapshot`
where `id` = ?;

```
* first comment line is to define the query name, use `name: ` mark. 
* also supports insert, update and delete, but they are generated by sql schema file.
### Generate
```bash
dalc --dialect=mysql \
     --out=example/generated/dal \
     --schema=example/generated/sqls/schema \
     --query=example/generated/sqls/query \
     --json_tags=true \
     --verbose=true
```
Command Args:
* dialect: sql dialect, can by mysql or postgres
* schema: sql schema file path or dir path
* query: sql query file path or dir path
* json_tags: enable to add json tag in table row struct and query result struct
* verbose: show verbose log
# Todo
[ ] postgres
## License

GNU GENERAL PUBLIC LICENSE 