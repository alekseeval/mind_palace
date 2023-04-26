CREATE OR REPLACE FUNCTION mind_palace_api.change_user (p_id int, p_name varchar = null, p_tg_id bigint = null)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    cur_user mind_palace.users;
    r_id int;
BEGIN
    SELECT * INTO cur_user
    FROM mind_palace.users
    WHERE id = p_id;

    if p_name is null then
        p_name = cur_user.name;
    end if;
    if p_tg_id is null then
        p_tg_id = cur_user.tg_id;
    end if;

    UPDATE mind_palace.users
    SET name=p_name, tg_id=p_tg_id
    WHERE id=p_id
    RETURNING id INTO r_id;

    RETURN r_id;
END;
$$;