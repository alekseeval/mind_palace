CREATE OR REPLACE FUNCTION mind_palace_api.create_user (p_name varchar, p_tg_id bigint)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
r_id int;
BEGIN
    INSERT INTO mind_palace.users(name, tg_id) VALUES (p_name, p_tg_id) RETURNING id INTO r_id;
    RETURN r_id;
END;
$$;