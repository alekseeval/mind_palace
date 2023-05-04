package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgresDB) CreateTheme(theme model.Theme) (*model.Theme, error) {
	row := p.db.QueryRow(`SELECT * FROM create_theme($1, $2, $3)`,
		theme.Title, theme.MainThemeId, theme.UserId)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	theme.Id = id
	return &theme, nil
}

func (p *PostgresDB) GetAllUserThemes(userId int) ([]*model.Theme, error) {
	rows, err := p.db.Queryx(`SELECT * FROM get_all_themes_for_user($1)`, userId)
	if err != nil {
		return nil, err
	}
	themes := make([]*model.Theme, 0)
	for rows.Next() {
		var t model.Theme
		err = rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
		themes = append(themes, &t)
	}
	return themes, nil
}

func (p *PostgresDB) ChangeTheme(theme *model.Theme) (*model.Theme, error) {
	row := p.db.QueryRowx(`SELECT * FROM create_theme($1, $2, $3, $4)`,
		theme.Id, theme.Title, theme.MainThemeId, theme.UserId)
	err := row.StructScan(theme)
	return theme, err
}

func (p *PostgresDB) DeleteTheme(id int) (int, error) {
	row := p.db.QueryRow(`SELECT * FROM delete_theme($1)`, id)
	var removeThemeId int
	err := row.Scan(&removeThemeId)
	return removeThemeId, err
}
