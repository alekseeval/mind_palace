create schema if not exists mind_palace;
create schema if not exists mind_palace_api;

ALTER ROLE CURRENT_ROLE SET SEARCH_PATH to mind_palace, mind_palace_api;

CREATE TABLE if not exists mind_palace.users (
    id serial PRIMARY KEY,
    name VARCHAR,
    tg_id bigint,

    UNIQUE(tg_id)
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
    title VARCHAR,

    UNIQUE(title)
);

CREATE TABLE if not exists mind_palace.notes (
    id serial PRIMARY KEY,
    title VARCHAR,
    text VARCHAR,
    note_type int,
    theme_id int,
    user_id int,

    UNIQUE(title, user_id),
    FOREIGN KEY (theme_id) REFERENCES mind_palace.themes (id),
    FOREIGN KEY (note_type) REFERENCES mind_palace.note_types (id),
    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id)
);