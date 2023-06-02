CREATE OR REPLACE FUNCTION mind_palace_api.change_note (p_id int, p_title varchar, p_text varchar, p_note_type int, p_theme_id int)
RETURNS notes
LANGUAGE plpgsql
AS
$$
DECLARE
    r_note notes;
    v_theme_id int;
    v_title varchar;
BEGIN
    select theme_id, title into v_theme_id, v_title from notes where id=p_id;
    if v_theme_id is null then
        RAISE SQLSTATE '80013' USING message = 'no such note';
    end if;

    if p_theme_id is not NULL AND not exists(SELECT * from themes where id=p_theme_id) then
        RAISE SQLSTATE '80007' USING message = 'no such theme';
    end if;

    if p_note_type is not null AND not EXISTS(SELECT * from note_types where id=p_note_type) then
        RAISE SQLSTATE '80011' USING message = 'wrong note type provided';
    end if;

    UPDATE mind_palace.notes SET
                                 title=coalesce(p_title, title),
                                 text=coalesce(p_text, text),
                                 note_type=coalesce(p_note_type, note_type),
                                 theme_id=coalesce(p_theme_id, theme_id)
    WHERE id=p_id RETURNING * INTO r_note;

    RETURN r_note;
END
$$;