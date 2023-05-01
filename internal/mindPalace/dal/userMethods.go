package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgreDB) SaveUser(user model.User) (*model.User, error) {
	return nil, nil
}

func (p *PostgreDB) GetUserByTgId(telegramIid int64) (*model.User, error) {
	return nil, nil
}

func (p *PostgreDB) ChangeUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (p *PostgreDB) DeleteUser(id int) (*model.User, error) {
	return nil, nil
}
