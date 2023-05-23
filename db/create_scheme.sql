create schema if not exists mind_palace;
create schema if not exists mind_palace_api;

ALTER ROLE CURRENT_ROLE SET SEARCH_PATH to mind_palace, mind_palace_api;

CREATE TABLE if not exists mind_palace.users (
    id serial PRIMARY KEY,
    name VARCHAR unique,
    tg_id bigint unique
);

CREATE TABLE if not exists mind_palace.themes (
    id serial PRIMARY KEY,
    title VARCHAR NOT NULL,
    main_theme_id int,
    user_id int,

    UNIQUE(title, user_id),
    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id),
    FOREIGN KEY (main_theme_id) REFERENCES mind_palace.themes (id)
);

CREATE TABLE if not exists mind_palace.note_types (
    id serial PRIMARY KEY,
    title VARCHAR NOT NULL ,

    UNIQUE(title)
);

CREATE TABLE if not exists mind_palace.notes (
    id serial PRIMARY KEY,
    title VARCHAR NOT NULL UNIQUE ,
    text VARCHAR NOT NULL,
    note_type int NOT NULL,
    theme_id int NOT NULL,

    FOREIGN KEY (theme_id) REFERENCES mind_palace.themes (id),
    FOREIGN KEY (note_type) REFERENCES mind_palace.note_types (id)
);

CREATE OR REPLACE VIEW v_theme AS SELECT t.id, t.title, t.main_theme_id, u.name as user_name FROM themes t LEFT JOIN users u ON t.user_id = u.id;