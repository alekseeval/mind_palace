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
        RAISE EXCEPTION 'no such note';
    end if;

    if not exists(SELECT * from themes where id=v_theme_id) then
        raise exception 'no such theme';
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