CREATE OR REPLACE FUNCTION mind_palace_api.change_note (p_id int, p_title varchar = null, p_text varchar = null, p_note_type int = null, p_theme_id int = null, p_user_id int = null)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    cur_note mind_palace.notes;
    r_id int;
BEGIN
    SELECT * INTO cur_note from mind_palace.notes WHERE id=p_id;

    if p_title is null then
        p_title = cur_note.title;
    end if;
    if p_text is null then
        p_text = cur_note.text;
    end if;
    if p_note_type is null then
        p_note_type = cur_note.note_type;
    end if;
    if p_theme_id is null then
        p_theme_id = cur_note.theme_id;
    end if;
    if p_user_id is null then
        p_user_id = cur_note.user_id;
    end if;

    UPDATE mind_palace.notes
    SET title=p_title, text=p_text, note_type=p_note_type, theme_id=p_theme_id, user_id=p_user_id
    WHERE id=p_id
    RETURNING id INTO r_id;

    RETURN r_id;
END
$$;