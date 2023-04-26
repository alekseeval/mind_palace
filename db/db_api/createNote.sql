CREATE OR REPLACE FUNCTION mind_palace_api.create_note (p_title varchar, p_text varchar, p_note_type int, p_theme_id int, p_user_id int)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    if NOT exists(SELECT * from mind_palace.users where id = p_user_id) then
        RAISE EXCEPTION 'user doesnt exists';
    end if;

    INSERT INTO mind_palace.notes(title, text, note_type, theme_id, user_id)
    VALUES (p_title, p_text, p_note_type, p_theme_id, p_user_id)
    RETURNING id INTO r_id;

    RETURN r_id;
END
$$;