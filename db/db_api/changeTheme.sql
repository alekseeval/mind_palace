CREATE OR REPLACE FUNCTION mind_palace_api.change_theme (p_theme_id int, p_title varchar = null, p_main_theme_id int = null)
RETURNS mind_palace.themes
LANGUAGE plpgsql
AS
$$
DECLARE
    changed_theme mind_palace.themes;
BEGIN
    SELECT * INTO changed_theme FROM mind_palace.themes WHERE id=p_theme_id;
    if p_title is null then
        p_title = changed_theme.title;
    end if;
    if p_main_theme_id is null then
        p_main_theme_id = changed_theme.main_theme_id;
    end if;
    if p_main_theme_id = 0 then
        p_main_theme_id = null;
    end if;
    UPDATE mind_palace.themes SET title=p_title, main_theme_id=p_main_theme_id WHERE id=p_theme_id RETURNING * INTO changed_theme;

    RETURN changed_theme;
END;
$$;