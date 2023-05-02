package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgresDB) CreateTheme(theme model.Theme) (*model.Theme, error) {
	return nil, nil
}

func (p *PostgresDB) GetAllUserThemes(user *model.User) ([]*model.Theme, error) {
	return nil, nil
}

func (p *PostgresDB) ChangeTheme(theme *model.Theme) (*model.Theme, error) {
	return nil, nil
}

func (p *PostgresDB) DeleteTheme(id int) (*model.Theme, error) {
	return nil, nil
}
