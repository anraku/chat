
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users (
	id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
	`name` varchar(64) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    udpated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deletion tinyint(1) NOT NULL DEFAULT 0
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;
