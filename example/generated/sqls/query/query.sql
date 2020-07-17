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

-- name: users_domain_events_list_v2
SELECT e.id, s.aggregate_id as `snapshot_aggregate_id`, e.aggregate_id as `event_aggregate_id`
FROM ddd_test.users_domain_events as e
         left join ddd_test.users_domain_snapshot as s on s.aggregate_id = e.aggregate_id
         left join ddd_test.users_domain_snapshot as s1 on s1.aggregate_id = e.aggregate_id
where e.id = ? and s.aggregate_id = ?;

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
(?,?,?,?,?,now(),'');

-- name: users_domain_events_update
UPDATE `ddd_test`.`users_domain_events`
SET
    `aggregate_name` = ?,
    `aggregate_id` = ?,
    `event_name` = '1',
    `event_id` = ?,
    `event_time` = now(),
    `event_data` = null
WHERE `id` = ?;

-- name: users_domain_events_delete
DELETE FROM `ddd_test`.`users_domain_events` WHERE `id` = ?;