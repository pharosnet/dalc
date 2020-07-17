-- name: users_domain_events_insert
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

-- name: users_domain_events_update
UPDATE `ddd_test`.`users_domain_events`
SET
`aggregate_name` = ?,
`aggregate_id` = ?,
`event_name` = ?,
`event_id` = ?,
`event_time` = ?,
`event_data` = ?
WHERE `id` = ?;

-- name: users_domain_events_delete
DELETE FROM `ddd_test`.`users_domain_events` WHERE `id` = ?;

-- name: users_domain_events_get
SELECT `users_domain_events`.`id`,
    `users_domain_events`.`aggregate_name`,
    `users_domain_events`.`aggregate_id`,
    `users_domain_events`.`event_name`,
    `users_domain_events`.`event_id`,
    `users_domain_events`.`event_time`,
    `users_domain_events`.`event_data`
FROM `ddd_test`.`users_domain_events`  WHERE `id` = ?;

-- name: users_domain_events_list
SELECT `users_domain_events`.`id`,
    `users_domain_events`.`aggregate_name`,
    `users_domain_events`.`aggregate_id`,
    `users_domain_events`.`event_name`,
    `users_domain_events`.`event_id`,
    `users_domain_events`.`event_time`,
    `users_domain_events`.`event_data`
FROM `ddd_test`.`users_domain_events`  WHERE `aggregate_id` = ? ORDER BY `id` DESC LIMIT ? OFFSET ?;

-- name: users_domain_events_list_v2
SELECT `users_domain_events`.`id`,
    `users_domain_events`.`aggregate_name`,
    `users_domain_events`.`aggregate_id`,
    `users_domain_events`.`event_name`,
    `users_domain_events`.`event_id`,
FROM `ddd_test`.`users_domain_events`  WHERE `aggregate_id` = ? ORDER BY `id` DESC LIMIT ? OFFSET ?;
