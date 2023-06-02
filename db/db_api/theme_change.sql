CREATE OR REPLACE FUNCTION mind_palace_api.change_theme (p_theme_id int, p_title varchar, p_main_theme_id int)
RETURNS v_theme
LANGUAGE plpgsql
AS
$$
DECLARE
    changed_theme v_theme;
    changed_theme_id int;
    v_main_theme v_theme;
    v_user_name varchar;
BEGIN
    if p_theme_id = p_main_theme_id then
        RAISE SQLSTATE '80005' USING message ='theme cant be main for itself';
    end if;

--  check main theme existence for user
    if p_main_theme_id is not null then
        SELECT * INTO v_main_theme from v_theme where id=p_main_theme_id;
        SELECT user_name INTO v_user_name from v_theme where id=p_theme_id;

        if v_main_theme is null then
            RAISE SQLSTATE '80004' USING message = 'no such main theme';
        end if;
        if v_main_theme.user_name != v_user_name then
            RAISE SQLSTATE '80003' USING message = 'main theme linked to other user';
        end if;
    end if;

--  do update
    UPDATE themes SET
                      title=coalesce(p_title, title),
                      main_theme_id=p_main_theme_id
    WHERE id=p_theme_id RETURNING id INTO changed_theme_id;
    if changed_theme_id is null then
        RAISE SQLSTATE '80007' USING message = 'no such theme';
    end if;

--  return view of changed theme
    SELECT * into changed_theme from v_theme where id=changed_theme_id;
    RETURN changed_theme;
END;
$$;