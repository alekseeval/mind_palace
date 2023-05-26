CREATE OR REPLACE FUNCTION mind_palace_api.create_note (p_title varchar, p_text varchar, p_note_type int, p_theme_id int)
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
    select id into p_theme_id from themes where id=p_theme_id;
    if p_theme_id is null then
        Raise exception 'no such theme';
    end if;

    if p_title is null then
        RAISE EXCEPTION 'no title provided';
    end if;
    if EXISTS(SELECT * from notes where theme_id=p_theme_id and title=p_title) then
        RAISE EXCEPTION 'note with title % already exists', p_title;
    end if;

    INSERT INTO notes(title, text, note_type, theme_id)
    VALUES (p_title, p_text, p_note_type, p_theme_id)
    RETURNING * INTO r_note;

    RETURN r_note;
END
$$;