package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgresDB) CreateNote(note model.Note) (*model.Note, error) {
	return nil, nil
}

func (p *PostgresDB) GetAllUserNotesByTheme(user *model.User, theme *model.Theme) ([]*model.Note, error) {
	return nil, nil
}

func (p *PostgresDB) ChangeNote(note *model.Note) (*model.Note, error) {
	return nil, nil
}

func (p *PostgresDB) DeleteNote(id int) (int, error) {
	return 0, nil
}
