-- +goose Up
-- +goose StatementBegin
CREATE TABLE credit_cards (
   id INT NOT NULL AUTO_INCREMENT,
   user_id INT NOT NULL,
   number varchar(255) DEFAULT NULL COMMENT 'number',
   created_at datetime DEFAULT NULL COMMENT 'creation date',
   updated_at datetime DEFAULT NULL COMMENT 'update date',
   deleted_at timestamp NULL DEFAULT NULL COMMENT 'delete date',
   PRIMARY KEY(id),
   CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='credit card';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE credit_cards;
-- +goose StatementEnd
