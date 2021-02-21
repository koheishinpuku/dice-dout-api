CREATE TABLE `users` (
  `id` CHAR NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` INT DEFAULT 0,
  `name` CHAR NOT NULL,
  `have_picies` TINYINT(5),
  `pice_prace` TINYINT(6),
  `gole_picies` TINYINT(2),
  UNIQUE (`id`, `name`),
  PRIMARY KEY (id)
);
CREATE TABLE `room` (
  `id` CHAR NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` INT DEFAULT 0,
  `name` CHAR NOT NULL,
  `player_one_id` CHAR,
  `player_two_id` CHAR,
  `player_three_id` CHAR,
  `player_four_id` CHAR,
  `room_status` TINYINT(5),
  UNIQUE (`id`, `name`),
  PRIMARY KEY (id)
);
ALTER TABLE `room`
ADD CONSTRAINT `room_ibfk_1` FOREIGN KEY (player_one_id) REFERENCES users(id);
ALTER TABLE `room`
ADD CONSTRAINT `room_ibfk_1` FOREIGN KEY (player_two_id) REFERENCES users(id);
ALTER TABLE `room`
ADD CONSTRAINT `room_ibfk_1` FOREIGN KEY (player_three_id) REFERENCES users(id);
ALTER TABLE `room`
ADD CONSTRAINT `room_ibfk_1` FOREIGN KEY (player_four_id) REFERENCES users(id);
