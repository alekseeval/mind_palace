CREATE OR REPLACE FUNCTION mind_palace_api.get_all_notes_by_theme(p_theme_id int)
RETURNS SETOF mind_palace.notes
LANGUAGE plpgsql
AS
$$
BEGIN
    if not exists(SELECT * from themes where id=p_theme_id) then
        RAISE SQLSTATE '80007' USING message = 'no such theme';
    end if;
    RETURN query
        SELECT * FROM mind_palace.notes WHERE theme_id=p_theme_id;
END
$$;