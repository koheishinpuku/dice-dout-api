-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
  `id` CHAR(30) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` INT DEFAULT 0,
  `name` CHAR(30) NOT NULL,
  `have_picies` TINYINT(5),
  `pice_place` TINYINT(6),
  `gole_picies` TINYINT(2),
  UNIQUE (`id`, `name`),
  PRIMARY KEY (id)
);
CREATE TABLE `rooms` (
  `id` CHAR(30) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` INT DEFAULT 0,
  `name` CHAR(30) NOT NULL,
  `player_one_id` CHAR(30),
  `player_two_id` CHAR(30),
  `player_three_id` CHAR(30),
  `player_four_id` CHAR(30),
  `room_status` TINYINT(5),
  UNIQUE (`id`, `name`),
  PRIMARY KEY (id)
);
ALTER TABLE `rooms`
ADD CONSTRAINT `room_ibfk_1` FOREIGN KEY (player_one_id) REFERENCES users(id);
ALTER TABLE `rooms`
ADD CONSTRAINT `room_ibfk_2` FOREIGN KEY (player_two_id) REFERENCES users(id);
ALTER TABLE `rooms`
ADD CONSTRAINT `room_ibfk_3` FOREIGN KEY (player_three_id) REFERENCES users(id);
ALTER TABLE `rooms`
ADD CONSTRAINT `room_ibfk_4` FOREIGN KEY (player_four_id) REFERENCES users(id);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `room` DROP FOREIGN KEY `room_ibfk_1`;
ALTER TABLE `room` DROP FOREIGN KEY `room_ibfk_2`;
ALTER TABLE `room` DROP FOREIGN KEY `room_ibfk_3`;
ALTER TABLE `room` DROP FOREIGN KEY `room_ibfk_4`;
drop.TABLE.users;
drop.TABLE.rooms;
