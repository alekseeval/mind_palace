package dal

import "MindPalace/internal/mindPalace/model"

func (p *PostgresDB) CreateNote(note model.Note) (*model.Note, error) {
	row := p.db.QueryRow(`SELECT * FROM create_note($1, $2, $3, $4, $5)`,
		note.Title, note.Text, note.NoteTypeId, note.ThemeId, note.UserId)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	note.Id = id
	return &note, nil
}

func (p *PostgresDB) GetAllUserNotesByTheme(userId int, themeId int) ([]*model.Note, error) {
	rows, err := p.db.Queryx(`SELECT * FROM get_notes($1, $2)`, userId, themeId)
	if err != nil {
		return nil, err
	}
	allNotes := make([]*model.Note, 0)
	for rows.Next() {
		var n model.Note
		err = rows.StructScan(&n)
		if err != nil {
			return nil, err
		}
		allNotes = append(allNotes, &n)
	}
	return allNotes, nil
}

func (p *PostgresDB) ChangeNote(note *model.Note) (*model.Note, error) {
	row := p.db.QueryRowx(`SELECT * FROM change_note($1, $2, $3, $4, $5, $6)`,
		note.Id, note.Title, note.Title, note.NoteTypeId, note.ThemeId, note.UserId)
	err := row.StructScan(note)
	return note, err
}

func (p *PostgresDB) DeleteNote(noteId int) (int, error) {
	row := p.db.QueryRow(`SELECT * FROM delete_note($1)`, noteId)
	var deletedNoteId int
	err := row.Scan(&deletedNoteId)
	return deletedNoteId, err
}
