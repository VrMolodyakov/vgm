BEGIN;

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    user_password TEXT NOT NULL,
    user_name TEXT NOT NULL,
    user_email TEXT NOT NULL,
    create_at TIMESTAMP NOT  NULL
);

CREATE TABLE user_roles(
    id SERIAL PRIMARY KEY,
    user_id int REFERENCES users (user_id),
    role_id int REFERENCES roles (role_id)
);
CREATE TABLE roles(
    role_id SERIAL PRIMARY KEY,
    role_name TEXT
);

INSERT INTO public.roles(role_name) VALUES ('user');
INSERT INTO public.roles(role_name) VALUES ('admin');
INSERT INTO public.users(user_name,user_email,user_password) VALUES ('admin','admin@gmail','admin');

COMMIT;