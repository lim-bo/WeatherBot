CREATE TABLE IF NOT EXISTS preferences (
    id SERIAL PRIMARY KEY,
    city varchar(64) not null default 'Moscow'
);