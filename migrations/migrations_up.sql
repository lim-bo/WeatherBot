CREATE TABLE IF NOT EXISTS preferences (
    id integer PRIMARY KEY,
    city varchar(64) not null default 'Moscow',
    status integer not null default 0
);