-- +migrate Up
CREATE SCHEMA IF NOT EXISTS hersshel;

CREATE TABLE IF NOT EXISTS hersshel.category (
	id SMALLSERIAL,
	name TEXT NOT NULL,
    CONSTRAINT pk_category PRIMARY KEY (id),
	CONSTRAINT uq_category_name UNIQUE (name)
);

CREATE INDEX idx_category_name ON hersshel.category (name);

CREATE TABLE IF NOT EXISTS hersshel.feed (
	id SERIAL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    url TEXT NOT NULL,
    name TEXT NOT NULL,
    website TEXT,
    description TEXT,
    image TEXT,
    category_id SMALLINT DEFAULT NULL,
    CONSTRAINT pk_feed PRIMARY KEY (id),
	CONSTRAINT fk_feed_category
        FOREIGN KEY (category_id) REFERENCES hersshel.category (id)
        ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT uq_feed_url UNIQUE (url)
);

CREATE INDEX idx_feed_name ON hersshel.feed (name);

CREATE TABLE IF NOT EXISTS hersshel.item (
    id BIGSERIAL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    title TEXT NOT NULL,
    author TEXT,
    content BYTEA,
    link TEXT NOT NULL,
    read BOOLEAN NOT NULL DEFAULT FALSE,
    starred BOOLEAN NOT NULL DEFAULT FALSE,
    feed_id INTEGER NOT NULL,
    CONSTRAINT pk_item PRIMARY KEY (id),
    CONSTRAINT fk_item_feed 
        FOREIGN KEY (feed_id) REFERENCES hersshel.feed (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uq_item_link UNIQUE (feed_id, link)
);

CREATE INDEX idx_item_read ON hersshel.item (read);
CREATE INDEX idx_item_starred ON hersshel.item (starred);

-- +migrate Down
DROP INDEX idx_category_name;
DROP INDEX idx_feed_name;
DROP INDEX idx_item_read;
DROP INDEX idx_item_starred;

DROP TABLE hersshel.category CASCADE;
DROP TABLE hersshel.feed CASCADE;
DROP TABLE hersshel.item CASCADE;