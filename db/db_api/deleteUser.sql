CREATE OR REPLACE FUNCTION mind_palace_api.delete_user(p_user_name varchar)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    DELETE FROM mind_palace.users WHERE name=p_user_name RETURNING id INTO r_id;
    if r_id is NULL then
        RAISE EXCEPTION 'No such user';
    end if;
    RETURN r_id;
END;
$$;