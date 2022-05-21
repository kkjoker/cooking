-- +goose Up
-- +goose StatementBegin

CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '内部ID'
,   `name` varchar(255) NOT NULL COMMENT '名前'
,	`address` varchar(255) NOT NULL COMMENT '住所'
,	`created_at` datetime NULL COMMENT '作成日時'
,	`updated_at` datetime NULL COMMENT '更新日時'
,   `deleted_at` datetime NULL COMMENT '削除日時'
,    PRIMARY KEY(id)
) ENGINE=InnoDB AUTO_INCREMENT=1 COMMENT='ユーザ' DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE `users`;

-- +goose StatementEnd
