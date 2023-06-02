CREATE OR REPLACE FUNCTION mind_palace_api.create_theme (p_title varchar, p_main_theme_id int, p_user_name varchar)
RETURNS v_theme
LANGUAGE plpgsql
AS
$$
DECLARE
    r_theme v_theme;
    v_user_id int;
    v_created_theme_id int;
    v_main_theme themes;
BEGIN
--  determine user
    if p_user_name is null THEN
        v_user_id=1;
    else
        select id INTO v_user_id FROM users where name=p_user_name;
        if v_user_id is NULL then
            RAISE SQLSTATE '80002' USING message = 'no such user';
        end if;
    end if;

--  check main theme existence for user
    if p_main_theme_id is not null then
        SELECT * INTO v_main_theme from themes where id=p_main_theme_id;
        if v_main_theme is null then
            RAISE SQLSTATE '80004' USING message = 'no such main theme';
        end if;
        if v_main_theme.user_id != v_user_id then
            RAISE SQLSTATE '80003' USING message = 'main theme linked to other user';
        end if;
    end if;

    INSERT INTO mind_palace.themes(title, main_theme_id, user_id) VALUES (p_title, p_main_theme_id, v_user_id) RETURNING id INTO v_created_theme_id;
    SELECT * INTO r_theme from v_theme WHERE id=v_created_theme_id;
    RETURN r_theme;
END
$$;