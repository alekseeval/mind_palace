package dal

import (
	"MindPalace/internal/mindPalace/model"
)

func (p *PostgresDB) SaveUser(user model.User) (*model.User, error) {
	row := p.db.QueryRow(`SELECT * FROM create_user($1, $2)`, user.Name, user.TelegramId)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	user.Id = id
	return &user, nil
}

func (p *PostgresDB) GetUserByTgId(telegramIid int64) (*model.User, error) {
	row := p.db.QueryRowx(`SELECT * FROM get_user_by_tg_id($1)`, telegramIid)
	var user model.User
	err := row.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresDB) ChangeUser(user *model.User) (*model.User, error) {
	row := p.db.QueryRowx(`SELECT * FROM change_user($1, $2, $3)`, user.Id, user.Name, user.TelegramId)
	err := row.StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *PostgresDB) DeleteUser(id int) (int, error) {
	row := p.db.QueryRowx(`SELECT * FROM delete_user($1)`, id)
	var deletedUserId int
	err := row.Scan(&deletedUserId)
	return deletedUserId, err
}
