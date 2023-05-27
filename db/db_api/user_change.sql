CREATE OR REPLACE FUNCTION mind_palace_api.change_user (p_id int, p_name varchar, p_tg_id bigint)
RETURNS mind_palace.users
LANGUAGE plpgsql
AS
$$
DECLARE
    cur_user mind_palace.users;
BEGIN
    UPDATE mind_palace.users SET
        tg_id=COALESCE(p_tg_id, tg_id),
        name=COALESCE(p_name, name)
    WHERE id=p_id
    RETURNING * INTO cur_user;

    if cur_user is NULL then
        RAISE EXCEPTION 'No such user';
    end if;

    RETURN cur_user;
END;
$$;