CREATE OR REPLACE FUNCTION mind_palace_api.get_all_themes_for_user (p_user_id int)
RETURNS mind_palace.themes
LANGUAGE plpgsql
AS
$$
Declare
    r_themes mind_palace.themes;
BEGIN
    SELECT *
    INTO r_themes
    FROM mind_palace.themes
    WHERE user_id=p_user_id;
    return r_themes;
END;
$$;