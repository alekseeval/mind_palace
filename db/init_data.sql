INSERT INTO mind_palace.note_types(id, title) VALUES (1, 'note'),
                                                     (2, 'question'),
                                                     (3, 'task')
ON CONFLICT (id) DO UPDATE
    SET title = excluded.title;

DO
$$
DECLARE
    users_count int;
BEGIN
    SELECT count(*) INTO users_count from mind_palace.users;
    if users_count = 0 then
        PERFORM nextval('mind_palace.users_id_seq');
    end if;
    INSERT INTO mind_palace.users(id, name, tg_id) VALUES (1, 'system', null)
    ON CONFLICT (id) DO UPDATE
        SET name = excluded.name;
END
$$;