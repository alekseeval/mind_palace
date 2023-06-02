CREATE OR REPLACE FUNCTION mind_palace_api.create_note (p_title varchar, p_text varchar, p_note_type int, p_theme_id int)
RETURNS notes
LANGUAGE plpgsql
AS
$$
DECLARE
    r_note notes;
BEGIN
    if p_theme_id is null then
        RAISE SQLSTATE '80008' USING message = 'no theme provided to function';
    end if;
    if not EXISTS(select id from themes where id=p_theme_id) then
        RAISE SQLSTATE '80007' USING message = 'no such theme exists';
    end if;
    if p_title is null then
        RAISE SQLSTATE '80009' USING message = 'empty title provided for create note';
    end if;
    if p_text is null then
        RAISE SQLSTATE '80012' USING message = 'empty text provided for create note';
    end if;
    if not EXISTS(SELECT * from note_types where id=p_note_type) then
        RAISE SQLSTATE '80011' USING message = 'wrong note type provided';
    end if;

    INSERT INTO notes(title, text, note_type, theme_id)
    VALUES (p_title, p_text, p_note_type, p_theme_id)
    RETURNING * INTO r_note;

    RETURN r_note;
END
$$;