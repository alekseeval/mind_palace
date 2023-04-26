CREATE OR REPLACE FUNCTION mind_palace_api.get_user_by_tg_id (p_tg_id bigint)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
    r_user mind_palace.users;
BEGIN
    SELECT * INTO r_user FROM mind_palace.users WHERE tg_id=p_tg_id;
    RETURN r_user;
END;
$$;