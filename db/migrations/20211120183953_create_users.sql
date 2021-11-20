-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
   id INT NOT NULL AUTO_INCREMENT,
   name varchar(255) DEFAULT NULL COMMENT 'username',
   age varchar(255) DEFAULT NULL COMMENT 'age',
   created_at datetime DEFAULT NULL COMMENT 'creation date',
   updated_at datetime DEFAULT NULL COMMENT 'update time',
   deleted_at timestamp NULL DEFAULT NULL COMMENT 'delete time',
   INDEX user_id (id),
   PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
