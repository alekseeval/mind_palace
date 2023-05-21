CREATE OR REPLACE FUNCTION mind_palace_api.create_theme (p_title varchar, p_main_theme_id int, p_user_name varchar)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
    v_user_id int;
BEGIN
    if p_user_name is null THEN
        v_user_id=1;
    else
        select id INTO v_user_id FROM users where name=p_user_name;
        if v_user_id is NULL then
            RAISE EXCEPTION 'no such user';
        end if;
    end if;

    if p_main_theme_id is not null then
        select id into p_main_theme_id from themes where id=p_main_theme_id and user_id=v_user_id;
        if p_main_theme_id is null then
            RAISE EXCEPTION 'no such main theme';
        end if;
    end if;

    INSERT INTO mind_palace.themes(title, main_theme_id, user_id) VALUES (p_title, p_main_theme_id, v_user_id) RETURNING id INTO r_id;
    RETURN r_id;
END;
$$;