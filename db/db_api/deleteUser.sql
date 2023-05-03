CREATE OR REPLACE FUNCTION mind_palace_api.delete_user (p_user_id int)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
    r_user mind_palace.users;
BEGIN
    DELETE FROM mind_palace.users WHERE id=p_user_id RETURNING * INTO r_user;
    RETURN r_user;
END;
$$;