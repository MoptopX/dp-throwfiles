-- +migrate Up

ALTER TABLE pageviews ALTER COLUMN timestamp TYPE TIMESTAMP WITHOUT TIME ZONE;

-- +migrate Down

ALTER TABLE pageviews ALTER COLUMN timestamp TYPE TIMESTAMP WITH TIME ZONE;