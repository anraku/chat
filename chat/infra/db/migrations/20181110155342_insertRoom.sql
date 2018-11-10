
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
insert into rooms (name, description) values ("test_room", "This is test room");


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

