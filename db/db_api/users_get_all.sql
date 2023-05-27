CREATE OR REPLACE FUNCTION mind_palace_api.get_users ()
RETURNS SETOF users
LANGUAGE plpgsql
AS
$$
BEGIN
    RETURN query select * from users;
END;
$$;