CREATE OR REPLACE FUNCTION mind_palace_api.get_notes (p_user_id int, p_theme_id int = null)
RETURNS SETOF mind_palace.notes
LANGUAGE plpgsql
AS
$$
BEGIN
    if NOT EXISTS(SELECT * FROM mind_palace.users where id=p_user_id) then
        RAISE EXCEPTION 'user doesnt exists';
    end if;

    if p_theme_id is null then
        RETURN query
            SELECT * FROM mind_palace.notes WHERE user_id=p_user_id;
    end if;
    RETURN query
        SELECT * FROM mind_palace.notes WHERE theme_id=p_theme_id and user_id=p_user_id;
END
$$;