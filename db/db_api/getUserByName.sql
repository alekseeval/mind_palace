CREATE OR REPLACE FUNCTION mind_palace_api.get_user_by_name (p_user_name varchar)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
    r_user mind_palace.users;
    v_cnt int;
BEGIN
    SELECT * INTO r_user FROM mind_palace.users WHERE name=p_user_name;
    GET DIAGNOSTICS v_cnt := ROW_COUNT;
    IF v_cnt = 0 THEN
        RAISE EXCEPTION 'There is no user %', p_user_name;
    END IF;
    RETURN r_user;
END;
$$;