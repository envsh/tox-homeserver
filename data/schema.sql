CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE `device` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `contact_id` INTEGER NULL, `uuid` TEXT NULL, `created` TEXT NULL, `updated` TEXT NULL);
CREATE UNIQUE INDEX `UQE_device_uuid` ON `device` (`uuid`);
CREATE INDEX `IDX_device_contact_id` ON `device` (`contact_id`);
CREATE TABLE `idgen` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
);
CREATE TABLE `message` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`updated`	TEXT,
	`created`	TEXT,
	`content`	TEXT,
	`mtype`	INTEGER,
	`contact_id`	INTEGER,
	`room_id`	INTEGER,
	`event_id`	INTEGER UNIQUE
, `sent` INTEGER NULL);
CREATE INDEX `IDX_message_contact_id` ON `message` (`contact_id`);
CREATE INDEX `IDX_message_room_id` ON `message` (`room_id`);
CREATE UNIQUE INDEX `UQE_message_event_id` ON `message` (`event_id`);
CREATE TABLE `sync_info` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `ct_id` INTEGER NULL, `next_batch` INTEGER NULL, `prev_batch` INTEGER NULL, `updated` TEXT NULL);
CREATE UNIQUE INDEX `UQE_sync_info_ct_id` ON `sync_info` (`ct_id`);
CREATE TABLE IF NOT EXISTS "contact" (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`pubkey`	TEXT NOT NULL UNIQUE,
	`name`	TEXT,
	`stmsg`	TEXT,
	`last_seen`	TEXT,
	`status`	INTEGER,
	`conn_status`	INTEGER,
	`created`	TEXT,
	`updated`	TEXT,
	`is_group`	INTEGER,
	`rt_id`	INTEGER,
	`is_peer`	INTEGER,
	`chat_type`	INTEGER,
	`is_friend`	INTEGER
);
CREATE UNIQUE INDEX `UQE_sync_info_siu` ON `sync_info` (`ct_id`,`next_batch`,`prev_batch`);
CREATE TABLE `setting` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `name` TEXT NULL, `value` TEXT NULL, `created` TEXT NULL, `updated` TEXT NULL);
CREATE UNIQUE INDEX `UQE_setting_name` ON `setting` (`name`);
CREATE INDEX `IDX_message_sent` ON `message` (`sent`);
CREATE INDEX `IDX_contact_is_peer` ON `contact` (`is_peer`);
CREATE UNIQUE INDEX `UQE_contact_pubkey` ON `contact` (`pubkey`);
CREATE INDEX `IDX_contact_name` ON `contact` (`name`);
CREATE INDEX `IDX_contact_is_group` ON `contact` (`is_group`);
CREATE INDEX `IDX_contact_is_friend` ON `contact` (`is_friend`);
