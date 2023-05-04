CREATE OR REPLACE FUNCTION mind_palace_api.change_theme (p_theme_id int, p_title varchar = null, p_main_theme_id int = null)
RETURNS mind_palace.themes
LANGUAGE plpgsql
AS
$$
DECLARE
    changed_theme mind_palace.themes;
BEGIN
    UPDATE mind_palace.themes
    SET title=p_title, main_theme_id=p_main_theme_id
    WHERE id=p_theme_id
    RETURNING * INTO changed_theme;
    RETURN changed_theme;
END;
$$;