CREATE OR REPLACE FUNCTION mind_palace_api.change_theme (p_theme_id int, p_title varchar, p_main_theme_id int)
RETURNS v_theme
LANGUAGE plpgsql
AS
$$
DECLARE
    changed_theme v_theme;
    changed_theme_id int;
BEGIN
    if p_theme_id = p_main_theme_id then
        RAISE 'theme cant be main for itself';
    end if;
--     TODO: можно ли тут обойтись без проверки и ловить стандартную ошибку при редактировании?
    if p_main_theme_id is not null then
        select id into p_main_theme_id from themes where id=p_main_theme_id;
        if p_main_theme_id is null then
            RAISE EXCEPTION 'no such main theme';
        end if;
    end if;

    UPDATE themes SET
                      title=coalesce(p_title, title),
                      main_theme_id=p_main_theme_id
    WHERE id=p_theme_id RETURNING id INTO changed_theme_id;
    if changed_theme_id is NULL then
        RAISE EXCEPTION 'no such theme';
    end if;

-- TODO: можно ли возвращать сразу запрос а не через переменную?
    SELECT * into changed_theme from v_theme where id=changed_theme_id;
    RETURN changed_theme;
END;
$$;