create schema mind_palace;
create schema mind_palace_api;

ALTER ROLE CURRENT_ROLE SET SEARCH_PATH to mind_palace, mind_palace_api;

CREATE TABLE if not exists mind_palace.users (
    id serial PRIMARY KEY,
    tg_id VARCHAR UNIQUE
);

CREATE TABLE if not exists mind_palace.themes (
    id serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    user_id int,

    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id)
);

CREATE TABLE if not exists mind_palace.ideas (
    id serial PRIMARY KEY,
    text VARCHAR,
    theme_id int,
    user_id int,

    FOREIGN KEY (theme_id) REFERENCES mind_palace.themes (id),
    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id)
);

CREATE TABLE if not exists mind_palace.questions (
    id serial PRIMARY KEY,
    text VARCHAR,
    main_idea int,
    theme_id int,
    number_of_correct_answers int,
    user_id int,

    FOREIGN KEY (theme_id) REFERENCES mind_palace.themes (id),
    FOREIGN KEY (main_idea) REFERENCES mind_palace.ideas (id),
    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id)
);