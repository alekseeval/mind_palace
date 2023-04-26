CREATE OR REPLACE FUNCTION mind_palace_api.create_theme (p_title varchar, p_main_theme_id int, p_user_id int)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    if p_user_id is null THEN
        p_user_id=1;
    end if;
    INSERT INTO mind_palace.themes(title, main_theme_id, user_id) VALUES (p_title, p_main_theme_id, p_user_id) RETURNING id INTO r_id;
    RETURN r_id;
END;
$$;