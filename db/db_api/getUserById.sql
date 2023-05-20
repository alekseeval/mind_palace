CREATE OR REPLACE FUNCTION mind_palace_api.get_user_by_id (p_id int)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
    r_user mind_palace.users;
    v_cnt int;
BEGIN
    SELECT * INTO r_user FROM mind_palace.users WHERE id=p_id;
    GET DIAGNOSTICS v_cnt := ROW_COUNT;
    IF v_cnt = 0 THEN
        RAISE EXCEPTION 'There is no user %', p_id;
    END IF;
    RETURN r_user;
END;
$$;