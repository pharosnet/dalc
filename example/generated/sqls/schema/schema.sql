USE `applications`;

-- name: business_group
CREATE TABLE `applications`.`business_group`
(
    `id`          varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL,
    `create_by`   varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL,
    `create_at`   datetime(6)                                            NOT NULL,
    `modify_by`   varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL,
    `modify_at`   datetime(6)                                            NOT NULL,
    `delete_by`   varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL,
    `delete_at`   datetime(6)                                            NOT NULL,
    `version`     bigint                                                 NOT NULL,
    `code`        varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL,
    `name`        varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL,
    `description` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `business_group_idx_biz_code` (`code`) USING BTREE,
    KEY `business_group_idx_biz_name` (`name`) USING BTREE,
    KEY `business_group_idx_create` (`create_by`, `create_at`) USING BTREE,
    KEY `business_group_idx_modify` (`modify_by`, `modify_at`) USING BTREE,
    KEY `business_group_idx_delete_by` (`delete_by`) USING BTREE,
    KEY `business_group_idx_delete_at` (`delete_at`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

-- name: business_code
CREATE TABLE `applications`.`business_code`
(
    `id`             varchar(63) COLLATE utf8mb4_bin  NOT NULL,
    `create_by`      varchar(63) COLLATE utf8mb4_bin  NOT NULL,
    `create_at`      datetime(6)                      NOT NULL,
    `modify_by`      varchar(63) COLLATE utf8mb4_bin  NOT NULL,
    `modify_at`      datetime(6)                      NOT NULL,
    `delete_by`      varchar(63) COLLATE utf8mb4_bin  NOT NULL,
    `delete_at`      datetime(6)                      NOT NULL,
    `version`        bigint                           NOT NULL,
    `code`  varchar(63) COLLATE utf8mb4_bin  NOT NULL,
    `group` varchar(63) COLLATE utf8mb4_bin  NOT NULL,
    `description`    varchar(512) COLLATE utf8mb4_bin NOT NULL,
    `text_code`      varchar(63) COLLATE utf8mb4_bin  NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `business_code_idx_biz_code_group` (`code`, `group`) USING BTREE,
    KEY `business_code_idx_create` (`create_by`, `create_at`) USING BTREE,
    KEY `business_code_idx_modify` (`modify_by`, `modify_at`) USING BTREE,
    KEY `business_code_idx_delete_by` (`delete_by`) USING BTREE,
    KEY `business_code_idx_delete_at` (`delete_at`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;
