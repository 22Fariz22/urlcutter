DROP TABLE IF EXISTS urls CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE urls
(
    url_id    UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    long      VARCHAR(350)              NOT NULL,
    short     VARCHAR(50)               NOT NULL,
    UNIQUE(long)
);

