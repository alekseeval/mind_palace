INSERT INTO mind_palace.note_types(id, title) VALUES (1, 'note'),
                                                     (2, 'question'),
                                                     (3, 'task')
ON CONFLICT (id) DO UPDATE
    SET title = excluded.title;

INSERT INTO mind_palace.users(id, name, tg_id) VALUES (1, 'system', null)
ON CONFLICT (id) DO UPDATE
    SET name = excluded.name;