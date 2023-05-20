CREATE OR REPLACE FUNCTION mind_palace_api.get_user_by_tg_id (p_tg_id bigint)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
    r_user mind_palace.users;
    v_cnt int;
BEGIN
    SELECT * INTO r_user FROM mind_palace.users WHERE tg_id=p_tg_id;
    GET DIAGNOSTICS v_cnt := ROW_COUNT;
    IF v_cnt = 0 THEN
        RAISE EXCEPTION 'There is no user with tg id=%', p_tg_id;
    END IF;
    RETURN r_user;
END;
$$;