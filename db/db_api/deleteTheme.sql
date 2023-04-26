CREATE OR REPLACE FUNCTION mind_palace_api.delete_theme (p_theme_id int)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    DELETE FROM mind_palace.themes WHERE id=p_theme_id RETURNING id INTO r_id;

    RETURN r_id;
END;
$$;