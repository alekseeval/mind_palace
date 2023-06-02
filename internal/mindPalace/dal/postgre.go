package dal

// TODO: Additional dbconn settings (maxconns, timeout)

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// PostgresDB is struct which implements model.IDAO interface and provides access to PostgresSQL DB
type PostgresDB struct {
	db     *sqlx.DB
	logger *logrus.Entry
}

// NewPostgresDB initialize PostgresDB struct
// error can be occurred by initial ping to db
func NewPostgresDB(config *configuration.Config, logger *logrus.Entry) (*PostgresDB, error) {
	dbConfig := config.System.DB
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s connect_timeout=%d",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName, dbConfig.Timeout)
	db, err := sqlx.Connect("postgres", connStr)
	return &PostgresDB{
		db:     db,
		logger: logger,
	}, err
}

// ---------------------------------------------------------------------------------------------------------------------
//  NOTE IDAO IMPLEMENTATION
// ---------------------------------------------------------------------------------------------------------------------

func (p *PostgresDB) SaveNote(note model.Note) (*model.Note, error) {
	queryRow := `SELECT * FROM create_note($1, $2, $3, $4)`
	queryParams := []interface{}{note.Title, note.Text, note.NoteTypeId, note.ThemeId}
	row := p.db.QueryRowx(queryRow, queryParams...)
	var dbNote model.Note
	err := row.StructScan(&dbNote)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	if err != nil {
		return nil, err
	}
	return &dbNote, nil
}

func (p *PostgresDB) GetAllNotesByTheme(themeId int) ([]*model.Note, error) {
	queryRow := `SELECT * FROM get_all_notes_by_theme($1)`
	queryParams := []interface{}{themeId}
	rows, err := p.db.Queryx(queryRow, queryParams...)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
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
	queryRow := `SELECT * FROM change_note($1, $2, $3, $4, $5)`
	queryParams := []interface{}{note.Id, note.Title, note.Text, note.NoteTypeId, note.ThemeId}
	row := p.db.QueryRowx(queryRow, queryParams...)
	err := row.StructScan(note)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	return note, err
}

func (p *PostgresDB) DeleteNote(noteId int) error {
	queryRow := `SELECT * FROM delete_note($1)`
	queryParams := []interface{}{noteId}
	_, err := p.db.Exec(queryRow, queryParams...)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	return err
}

// ---------------------------------------------------------------------------------------------------------------------
//  THEME IDAO IMPLEMENTATION
// ---------------------------------------------------------------------------------------------------------------------

func (p *PostgresDB) SaveTheme(theme model.Theme) (*model.Theme, error) {
	queryRow := `SELECT * FROM create_theme($1, $2, $3)`
	queryParams := []interface{}{theme.Title, theme.MainThemeId, theme.UserName}
	row := p.db.QueryRowx(queryRow, queryParams...)
	var dbTheme model.Theme
	err := row.StructScan(&dbTheme)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	if err != nil {
		return nil, err
	}
	return &dbTheme, nil
}

func (p *PostgresDB) GetAllUserThemes(userName *string) ([]*model.Theme, error) {
	queryRow := `SELECT * FROM get_all_themes_for_user($1)`
	queryParams := []interface{}{userName}
	rows, err := p.db.Queryx(queryRow, queryParams...)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
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
	queryRow := `SELECT * FROM change_theme($1, $2, $3)`
	queryParams := []interface{}{theme.Id, theme.Title, theme.MainThemeId}
	row := p.db.QueryRowx(queryRow, queryParams...)
	err := row.StructScan(theme)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	return theme, err
}

func (p *PostgresDB) DeleteTheme(themeId int) error {
	queryRow := `SELECT * FROM delete_theme($1)`
	queryParams := []interface{}{themeId}
	_, err := p.db.Exec(queryRow, queryParams...)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	return err
}

// ---------------------------------------------------------------------------------------------------------------------
//  USER IDAO IMPLEMENTATION
// ---------------------------------------------------------------------------------------------------------------------

func (p *PostgresDB) SaveUser(user model.User) (*model.User, error) {
	queryRow := `SELECT * FROM create_user($1, $2)`
	queryParams := []interface{}{user.Name, user.TelegramId}
	row := p.db.QueryRowx(queryRow, queryParams...)
	var dbUser model.User
	err := row.StructScan(&dbUser)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	if err != nil {
		return nil, err
	}
	return &dbUser, nil
}

func (p *PostgresDB) GetUserByTgId(telegramId int64) (*model.User, error) {
	queryRow := `SELECT * FROM get_user_by_tg_id($1)`
	queryParams := []interface{}{telegramId}
	row := p.db.QueryRowx(queryRow, queryParams...)
	var user model.User
	err := row.StructScan(&user)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresDB) GetUserById(userId int) (*model.User, error) {
	queryRow := `SELECT * FROM get_user_by_id($1)`
	queryParams := []interface{}{userId}
	row := p.db.QueryRowx(queryRow, queryParams...)
	var user model.User
	err := row.StructScan(&user)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresDB) ChangeUser(user *model.User) (*model.User, error) {
	queryRow := `SELECT * FROM change_user($1, $2, $3)`
	queryParams := []interface{}{user.Id, user.Name, user.TelegramId}
	row := p.db.QueryRowx(queryRow, queryParams...)
	err := row.StructScan(user)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *PostgresDB) DeleteUser(userId int) error {
	queryRow := `SELECT * FROM delete_user($1)`
	queryParams := []interface{}{userId}
	_, err := p.db.Exec(queryRow, queryParams...)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	return err
}

func (p *PostgresDB) GetAllUsers() ([]*model.User, error) {
	queryRow := `SELECT * FROM get_users()`
	queryParams := []interface{}{}
	rows, err := p.db.Queryx(queryRow)
	p.logger.WithFields(logrus.Fields{
		"params": queryParams,
		"query":  queryRow,
	}).Info("db query")
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0)
	for rows.Next() {
		var u model.User
		if err = rows.StructScan(&u); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
