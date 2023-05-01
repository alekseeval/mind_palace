package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgreDB) CreateNote(note model.Note) (*model.Note, error) {
	return nil, nil
}

func (p *PostgreDB) GetAllUserNotesByTheme(user *model.User, theme *model.Theme) ([]*model.Note, error) {
	return nil, nil
}

func (p *PostgreDB) ChangeNote(note *model.Note) (*model.Note, error) {
	return nil, nil
}

func (p *PostgreDB) DeleteNote(id int) (*model.Note, error) {
	return nil, nil
}
