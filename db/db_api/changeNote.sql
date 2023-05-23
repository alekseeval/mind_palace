CREATE OR REPLACE FUNCTION mind_palace_api.change_note (p_id int, p_title varchar, p_text varchar, p_note_type int, p_theme_id int)
RETURNS notes
LANGUAGE plpgsql
AS
$$
DECLARE
    r_note notes;
BEGIN

    if p_theme_id is null then
        RAISE EXCEPTION 'no theme provided';
    end if;
    if not exists(SELECT * from themes where id=p_theme_id) then
        raise exception 'no such theme';
    end if;

    if p_title is null then
        raise exception 'no title provided';
    end if;
    if exists(SELECT * from notes where theme_id=p_theme_id and title=p_title) then
        raise exception 'note with title % already exists', p_title;
    end if;

    UPDATE mind_palace.notes
    SET title=p_title, text=p_text, note_type=p_note_type, theme_id=p_theme_id
    WHERE id=p_id
    RETURNING * INTO r_note;

    if r_note is null then
        RAISE EXCEPTION 'no such note';
    end if;

    RETURN r_note;
END
$$;