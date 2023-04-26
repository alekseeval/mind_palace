INSERT INTO mind_palace.note_types(id, title) VALUES (1, 'note'),
                                                     (2, 'question'),
                                                     (3, 'task')
ON CONFLICT (id) DO UPDATE
    SET title = excluded.title;