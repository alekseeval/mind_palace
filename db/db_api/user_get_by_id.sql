CREATE OR REPLACE FUNCTION mind_palace_api.get_user_by_id (p_id int)
RETURNS users
LANGUAGE plpgsql
AS
$$
DECLARE
    r_user users;
    v_cnt int;
BEGIN
    SELECT * INTO r_user FROM mind_palace.users WHERE id=p_id;
    GET DIAGNOSTICS v_cnt := ROW_COUNT;
    IF v_cnt = 0 THEN
        RAISE SQLSTATE '80002' USING message = 'no such user';
    END IF;
    RETURN r_user;
END;
$$;