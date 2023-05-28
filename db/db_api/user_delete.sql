CREATE OR REPLACE FUNCTION mind_palace_api.delete_user(p_id int)
RETURNS void
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    DELETE FROM mind_palace.users WHERE id=p_id RETURNING id INTO r_id;
    if r_id is NULL then
        RAISE SQLSTATE '80002' USING MESSAGE = 'no such user';
    end if;
END
$$;