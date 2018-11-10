
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table rooms (
	id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
	name varchar(64),
	description varchar(255)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table rooms;
