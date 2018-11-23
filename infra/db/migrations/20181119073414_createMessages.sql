
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table messages (
	id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id integer(11) NOT NULL,
    room_id integer(11) NOT NULL,
	message varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    udpated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deletion tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table messages;
