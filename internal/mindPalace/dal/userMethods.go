package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgresDB) SaveUser(user model.User) (*model.User, error) {
	return nil, nil
}

func (p *PostgresDB) GetUserByTgId(telegramIid int64) (*model.User, error) {
	return nil, nil
}

func (p *PostgresDB) ChangeUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (p *PostgresDB) DeleteUser(id int) (*model.User, error) {
	return nil, nil
}
