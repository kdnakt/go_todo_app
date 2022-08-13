CREATE TABLE `user`
(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザー識別子',
    `name`      varchar(20) NOT NULL COMMENT 'ユーザー名',
    `password`  varchar(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role`      varchar(80) NOT NULL COMMENT 'ロール',
    `created`   DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified`  DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `task`
(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスク識別子',
    `title`     varchar(128) NOT NULL COMMENT 'タスクのタイトル',
    `status`    varchar(20)  NOT NULL COMMENT 'タスクの状態',
    `created`   DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified`  DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';
