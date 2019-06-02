CREATE DATABASE chat DEFAULT CHARACTER SET utf8;
use chat


create table rooms (
	id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
	name varchar(64) NOT NULL,
	description varchar(255),
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    udpated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deletion tinyint(1) NOT NULL DEFAULT 0
);

create table users (
	id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
	`name` varchar(64) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    udpated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deletion tinyint(1) NOT NULL DEFAULT 0
);

create table messages (
	id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_name varchar(128) NOT NULL,
    room_id integer(11) NOT NULL,
	message varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    udpated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deletion tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
