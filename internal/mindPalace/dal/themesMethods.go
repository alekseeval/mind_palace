package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgreDB) CreateTheme(theme model.Theme) (*model.Theme, error) {
	return nil, nil
}

func (p *PostgreDB) GetAllUserThemes(user *model.User) ([]*model.Theme, error) {
	return nil, nil
}

func (p *PostgreDB) ChangeTheme(theme *model.Theme) (*model.Theme, error) {
	return nil, nil
}

func (p *PostgreDB) DeleteTheme(id int) (*model.Theme, error) {
	return nil, nil
}
