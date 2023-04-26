CREATE OR REPLACE FUNCTION mind_palace_api.delete_note (p_id int)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    DELETE FROM mind_palace.notes WHERE id=p_id RETURNING id INTO r_id;
    RETURN r_id;
END
$$;