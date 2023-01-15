BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE currency
(
    currency_id SERIAL PRIMARY KEY,
    name   TEXT,
    symbol TEXT
);


CREATE TABLE album
(
    album_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    create_at DATE NOT NULL
);


CREATE TABLE album_info
(
    album_info_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    catalog_number TEXT,
    image_srs TEXT,
    barcode TEXT,
    release_date DATE NOT NULL,
    price NUMERIC(5,2) NOT NULL,
    currency_id INT REFERENCES currency (currency_id),
    media_format TEXT,
    classification TEXT,
    publisher TEXT

);

CREATE TABLE musical_profession
(
    profession_id SERIAL PRIMARY KEY,
    professional_title TEXT
);

CREATE TABLE person
(
    person_id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

create table credit(
    person_id INT REFERENCES person (person_id),
    album_id INT REFERENCES music_album (album_id),
    profession_id INT REFERENCES musical_profession (profession_id),
    PRIMARY KEY(person_id,album_id,profession_id)

); 

create table tracklist 
(
    track_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL
); 
