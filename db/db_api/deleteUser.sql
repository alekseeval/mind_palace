CREATE OR REPLACE FUNCTION mind_palace_api.delete_user(p_id int)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    DELETE FROM mind_palace.users WHERE id=p_id RETURNING id INTO r_id;
    if r_id is NULL then
        RAISE EXCEPTION 'No such user';
    end if;
    RETURN r_id;
END;
$$;