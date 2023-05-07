CREATE OR REPLACE FUNCTION mind_palace_api.get_all_themes_for_user (p_user_id int)
RETURNS SETOF mind_palace.themes
LANGUAGE plpgsql
AS
$$
BEGIN
    return query
        SELECT * FROM mind_palace.themes WHERE user_id=p_user_id;
END;
$$;