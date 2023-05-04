CREATE OR REPLACE FUNCTION mind_palace_api.change_user (p_id int, p_name varchar = null, p_tg_id bigint = null)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
    cur_user mind_palace.users;
BEGIN
    UPDATE mind_palace.users
    SET name=p_name, tg_id=p_tg_id
    WHERE id=p_id
    RETURNING * INTO cur_user;

    RETURN cur_user;
END;
$$;