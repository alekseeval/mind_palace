CREATE OR REPLACE FUNCTION mind_palace_api.get_user_by_id (p_id int)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
r_user mind_palace.users;
BEGIN
SELECT * INTO r_user FROM mind_palace.users WHERE id=p_id;
RETURN r_user;
END;
$$;