create schema mind_palace;
create schema mind_palace_api;

ALTER ROLE CURRENT_ROLE SET SEARCH_PATH to mind_palace, mind_palace_api;

CREATE TABLE if not exists mind_palace.users (
    id serial PRIMARY KEY,
    name VARCHAR,
    tg_id VARCHAR UNIQUE
);

CREATE TABLE if not exists mind_palace.themes (
    id serial PRIMARY KEY,
    title VARCHAR NOT NULL,

    main_theme_id int,
    user_id int,

    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id),
    FOREIGN KEY (main_theme_id) REFERENCES mind_palace.themes (id)
);

CREATE TABLE if not exists mind_palace.questions (
    id serial PRIMARY KEY,
    title VARCHAR,
    theme_id int,
    correct_answered_time int,

    user_id int,

    FOREIGN KEY (theme_id) REFERENCES mind_palace.themes (id),
    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id)
);

CREATE TABLE if not exists mind_palace.ideas (
    id serial PRIMARY KEY,
    title VARCHAR,
    text VARCHAR,

    is_answer boolean,
    question_id int,
    user_id int,

    FOREIGN KEY (question_id) REFERENCES mind_palace.themes (id),
    FOREIGN KEY (user_id) REFERENCES mind_palace.users (id)
);