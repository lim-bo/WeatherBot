CREATE TABLE IF NOT EXISTS preferences (
    user_id integer PRIMARY KEY,
    city varchar(64) not null default '',
    status integer not null default 0
);