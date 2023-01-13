BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE currency
(
    currency_id SERIAL PRIMARY KEY,
    name   TEXT,
    symbol TEXT
);

CREATE TABLE music_album
(
    album_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    catalog_number TEXT,
    barcode TEXT,
    release_date DATE NOT NULL,
    price NUMERIC(5,2) NOT NULL,
    currency_id INT REFERENCES currency (currency_id),
    media_format TEXT,
    classification TEXT,
    publisher TEXT

);

