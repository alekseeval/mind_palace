CREATE OR REPLACE FUNCTION mind_palace_api.delete_user (p_user_id int)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
BEGIN
    DELETE FROM mind_palace.users WHERE id=p_user_id RETURNING id INTO r_id;
    RETURN r_id;
END;
$$;