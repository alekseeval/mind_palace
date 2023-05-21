CREATE OR REPLACE FUNCTION mind_palace_api.change_theme (p_theme_id int, p_title varchar, p_main_theme_id int)
RETURNS mind_palace.v_theme
LANGUAGE plpgsql
AS
$$
DECLARE
    changed_theme mind_palace.v_theme;
    changed_theme_id int;
BEGIN
    if p_theme_id is NULL then
        RAISE exception 'no theme id provided';
    end if;

    if p_main_theme_id is not null then
        select id into p_main_theme_id from themes where id=p_main_theme_id;
        if p_main_theme_id is null then
            RAISE EXCEPTION 'no such main theme';
        end if;
    end if;

    UPDATE mind_palace.themes
    SET title=p_title, main_theme_id=p_main_theme_id
    WHERE id=p_theme_id
    RETURNING id INTO changed_theme_id;
    if changed_theme_id is NULL then
        RAISE EXCEPTION 'no such theme';
    end if;

    SELECT * into changed_theme from v_theme where id=changed_theme_id;
    RETURN changed_theme;
END;
$$;