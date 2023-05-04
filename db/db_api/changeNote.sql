CREATE OR REPLACE FUNCTION mind_palace_api.change_note (p_id int, p_title varchar = null, p_text varchar = null, p_note_type int = null, p_theme_id int = null, p_user_id int = null)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    UPDATE mind_palace.notes
    SET title=p_title, text=p_text, note_type=p_note_type, theme_id=p_theme_id, user_id=p_user_id
    WHERE id=p_id
    RETURNING id INTO r_id;

    RETURN r_id;
END
$$;