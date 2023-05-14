package dal

// TODO: Additional dbconn settings (maxconns, timeout)

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// PostgresDB is struct which implements model.IDAO interface and provides access to PostgresSQL DB
type PostgresDB struct {
	db *sqlx.DB
}

// NewPostgresDB initialize PostgresDB struct
// error can be occurred by initial ping to db
func NewPostgresDB(config *configuration.Config) (*PostgresDB, error) {
	dbConfig := config.System.DB
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s connect_timeout=%d",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName, dbConfig.Timeout)
	db, err := sqlx.Connect("postgres", connStr)
	return &PostgresDB{
		db: db,
	}, err
}

// ---------------------------------------------------------------------------------------------------------------------
//  NOTE IDAO IMPLEMENTATION
// ---------------------------------------------------------------------------------------------------------------------

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
	rows, err := p.db.Queryx(`SELECT * FROM get_all_notes_for_user_by_theme($1, $2)`, userId, themeId)
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

// ---------------------------------------------------------------------------------------------------------------------
//  THEME IDAO IMPLEMENTATION
// ---------------------------------------------------------------------------------------------------------------------

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
	row := p.db.QueryRowx(`SELECT * FROM change_theme($1, $2, $3)`,
		theme.Id, theme.Title, theme.MainThemeId)
	err := row.StructScan(theme)
	return theme, err
}

func (p *PostgresDB) DeleteTheme(themeId int) (int, error) {
	row := p.db.QueryRow(`SELECT * FROM delete_theme($1)`, themeId)
	var removeThemeId int
	err := row.Scan(&removeThemeId)
	return removeThemeId, err
}

// ---------------------------------------------------------------------------------------------------------------------
//  USER IDAO IMPLEMENTATION
// ---------------------------------------------------------------------------------------------------------------------

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

func (p *PostgresDB) GetUserByTgId(telegramId int64) (*model.User, error) {
	row := p.db.QueryRowx(`SELECT * FROM get_user_by_tg_id($1)`, telegramId)
	var user model.User
	err := row.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresDB) GetUserById(userId int) (*model.User, error) {
	row := p.db.QueryRowx(`SELECT * FROM get_user_by_id($1)`, userId)
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

func (p *PostgresDB) DeleteUser(userId int) (int, error) {
	row := p.db.QueryRowx(`SELECT * FROM delete_user($1)`, userId)
	var deletedUserId int
	err := row.Scan(&deletedUserId)
	return deletedUserId, err
}
