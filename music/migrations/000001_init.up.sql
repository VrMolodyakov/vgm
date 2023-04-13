BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE album
(
    album_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    released_at DATE NOT NULL,
    created_at DATE NOT NULL
);


CREATE TABLE album_info
(
    album_info_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    album_id UUID  REFERENCES album (album_id),
    catalog_number TEXT,
    full_image_srs TEXT,
    small_image_srs TEXT,
    barcode TEXT,
    price NUMERIC(8,2) NOT NULL,
    currency_code TEXT,
    media_format TEXT,
    classification TEXT,
    publisher TEXT

);

CREATE TABLE person
(
    person_id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    birth_date DATE
);

create table credit(
    credit_id SERIAL PRIMARY KEY,
    person_id INT REFERENCES person (person_id),
    album_id UUID  REFERENCES album (album_id),
    credit_role TEXT 
); 

create table track
(
    track_id SERIAL PRIMARY KEY,
    album_id UUID  REFERENCES album (album_id),
    title TEXT NOT NULL,
    duration TEXT NOT NULL
); 

COMMIT;