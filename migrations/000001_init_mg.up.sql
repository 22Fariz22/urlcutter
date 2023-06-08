DROP TABLE IF EXISTS urls CASCADE;


CREATE TABLE urls
(
    url_id    UUID PRIMARY KEY         NOT NULL ,
    long      VARCHAR(350)              NOT NULL,
    short     VARCHAR(50)               NOT NULL
);

