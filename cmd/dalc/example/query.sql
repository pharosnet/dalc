-- :w
-- name: CreateDomainEvent
-- result: id
INSERT INTO `ddd_test`.`users_domain_events`
(`id`,
`aggregate_name`,
`aggregate_id`,
`event_name`,
`event_id`,
`event_time`,
`event_data`)
VALUES
(?,?,?,?,?,?,?);

-- :w
-- name: ModifyDomainEvent
UPDATE `ddd_test`.`users_domain_events`
SET
`aggregate_name` = ?,
`aggregate_id` = ?,
`event_name` = ?,
`event_id` = ?,
`event_time` = ?,
`event_data` = ?
WHERE `id` = ?;

-- :w
-- name: DeleteDomainEvent
DELETE FROM `ddd_test`.`users_domain_events` WHERE `id` = ?;

-- :r
-- name: GetDomainEvent
-- ref: `ddd_test`.`users_domain_events`
SELECT `users_domain_events`.`id`,
    `users_domain_events`.`aggregate_name`,
    `users_domain_events`.`aggregate_id`,
    `users_domain_events`.`event_name`,
    `users_domain_events`.`event_id`,
    `users_domain_events`.`event_time`,
    `users_domain_events`.`event_data`
FROM `ddd_test`.`users_domain_events`  WHERE `id` = ?;

-- :r
-- name: ListDomainEvents
-- ref: `ddd_test`.`users_domain_events`
SELECT `users_domain_events`.`id`,
    `users_domain_events`.`aggregate_name`,
    `users_domain_events`.`aggregate_id`,
    `users_domain_events`.`event_name`,
    `users_domain_events`.`event_id`,
    `users_domain_events`.`event_time`,
    `users_domain_events`.`event_data`
FROM `ddd_test`.`users_domain_events`  WHERE `aggregate_id` = ? ORDER BY `id` DESC LIMIT ? OFFSET ?;

-- :r
-- name: ListDomainEventsXX
-- result:
-- aggregate_name sql.NullString
-- aggregate_id sql.NullString
-- event_name sql.NullString
-- event_id sql.NullString
SELECT `users_domain_events`.`id`,
    `users_domain_events`.`aggregate_name`,
    `users_domain_events`.`aggregate_id`,
    `users_domain_events`.`event_name`,
    `users_domain_events`.`event_id`,
FROM `ddd_test`.`users_domain_events`  WHERE `aggregate_id` = ? ORDER BY `id` DESC LIMIT ? OFFSET ?;
