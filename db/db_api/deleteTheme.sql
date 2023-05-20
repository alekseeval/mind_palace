CREATE OR REPLACE FUNCTION mind_palace_api.delete_theme (p_theme_id int)
RETURNS int
LANGUAGE plpgsql
AS
$$
DECLARE
    r_id int;
    v_sh_cnt int;
BEGIN
    if p_theme_id is NULL then
        RAISE EXCEPTION 'no theme id provided';
    end if;
    select count(*) into v_sh_cnt from themes where main_theme_id=p_theme_id;
    if v_sh_cnt != 0 then
        RAISE EXCEPTION 'theme have subthemes';
    end if;
    DELETE FROM mind_palace.themes WHERE id=p_theme_id RETURNING id INTO r_id;
    if r_id is NULL then
        RAISE EXCEPTION 'no such theme';
    end if;
    RETURN r_id;
END;
$$;