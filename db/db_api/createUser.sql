CREATE OR REPLACE FUNCTION mind_palace_api.create_user (p_name varchar, p_tg_id bigint)
RETURNS users
LANGUAGE plpgsql
AS
$$
DECLARE
    r_user users;
BEGIN
    if length(p_name) > 30 then
        RAISE SQLSTATE '80001' USING MESSAGE = 'user name too long';
    end if;
    INSERT INTO mind_palace.users(name, tg_id) VALUES (p_name, p_tg_id) RETURNING * INTO r_user;
    RETURN r_user;
END
$$;