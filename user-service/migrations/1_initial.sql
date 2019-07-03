-- +migrate Up
CREATE TABLE "users"
(
    "user_id" text NOT NULL,
    "phone" text NOT NULL,
    "password" text NOT NULL,
    "avatar" text NOT NULL,
    "display_name" text NOT NULL,
    "role" text NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id, phone),
    CONSTRAINT users_phone_key UNIQUE (phone),
    CONSTRAINT users_userid_key UNIQUE (user_id)
)

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;



