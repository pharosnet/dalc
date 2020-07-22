-- name: business_code_get_by_group_and_code
SELECT `business_code`.`id`,
       `business_code`.`create_by`,
       `business_code`.`create_at`,
       `business_code`.`modify_by`,
       `business_code`.`modify_at`,
       `business_code`.`delete_by`,
       `business_code`.`delete_at`,
       `business_code`.`version`,
       `business_code`.`code`,
       `business_code`.`group`,
       `business_code`.`description`,
       `business_code`.`text_code`
FROM `applications`.`business_code`
WHERE `group` = ?
  AND `code` = ?;

-- name: business_code_delete_by_code
DELETE
FROM `applications`.`business_code`
WHERE `group` = ?
  AND `code`;

-- name: business_code_list_by_group
SELECT `business_code`.`id`,
       `business_code`.`create_by`,
       `business_code`.`create_at`,
       `business_code`.`modify_by`,
       `business_code`.`modify_at`,
       `business_code`.`delete_by`,
       `business_code`.`delete_at`,
       `business_code`.`version`,
       `business_code`.`code`,
       `business_code`.`group`,
       `business_code`.`description`,
       `business_code`.`text_code`
FROM `applications`.`business_code`
WHERE `group` = ?
ORDER BY `code`
LIMIT ? OFFSET ?;

-- name: business_group_get_by_code
SELECT `business_group`.`id`,
       `business_group`.`create_by`,
       `business_group`.`create_at`,
       `business_group`.`modify_by`,
       `business_group`.`modify_at`,
       `business_group`.`delete_by`,
       `business_group`.`delete_at`,
       `business_group`.`version`,
       `business_group`.`code`,
       `business_group`.`name`,
       `business_group`.`description`
FROM `applications`.`business_group`
WHERE `code` = ?;

-- name: business_group_delete_by_code
DELETE
FROM `applications`.`business_group`
WHERE `code` = ?;

-- name: business_group_list
SELECT `business_group`.`id`,
       `business_group`.`create_by`,
       `business_group`.`create_at`,
       `business_group`.`modify_by`,
       `business_group`.`modify_at`,
       `business_group`.`delete_by`,
       `business_group`.`delete_at`,
       `business_group`.`version`,
       `business_group`.`code`,
       `business_group`.`name`,
       `business_group`.`description`
FROM `applications`.`business_group`
ORDER BY `code`
LIMIT ? OFFSET ?;