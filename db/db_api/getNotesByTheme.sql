CREATE OR REPLACE FUNCTION mind_palace_api.get_all_notes_by_theme(p_theme_id int)
RETURNS SETOF mind_palace.notes
LANGUAGE plpgsql
AS
$$
BEGIN
    if p_theme_id is null then
        raise exception 'no theme provided';
    end if;
    if not exists(SELECT * from themes where id=p_theme_id) then
        raise exception 'no such theme';
    end if;
    RETURN query
        SELECT * FROM mind_palace.notes WHERE theme_id=p_theme_id;
END
$$;