CREATE OR REPLACE FUNCTION mind_palace_api.delete_note (p_id int)
RETURNS void
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    DELETE FROM mind_palace.notes WHERE id=p_id RETURNING id INTO r_id;
    if r_id is null then
        RAISE SQLSTATE '80013' USING message = 'no such note';
    end if;
END
$$;