CREATE OR REPLACE FUNCTION mind_palace_api.get_all_themes_for_user (p_user_name varchar)
RETURNS SETOF mind_palace.v_theme
LANGUAGE plpgsql
AS
$$
DECLARE
    v_user_id int;
BEGIN
    IF p_user_name is NULL then
        return query
            SELECT * FROM mind_palace.v_theme;
        return;
    end if;
    select id into v_user_id from users where name=p_user_name;
    if v_user_id is null then
        RAISE EXCEPTION 'no such user %', p_user_name;
    end if;
    return query
        SELECT * FROM mind_palace.v_theme WHERE user_name=p_user_name;
END;
$$;