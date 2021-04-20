package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/tomashphill/exquisite-corpse-go/pkg/models"
)

type ExqCorpModelSqlite struct {
	*sql.DB
}

const createUserTable = `CREATE TABLE IF NOT EXISTS corpses (
	name VARCHAR(20) PRIMARY KEY
	stage TINYINT NOT NULL DEFAULT 1
)`

const createCorpseTable = `CREATE TABLE IF NOT EXISTS users (
	name VARCHAR(20) PRIMARY KEY
)`

// OpenDB will open a sqlite3 database, creating users and corpses tables if they do no already exist
func OpenDB(dsn string) (models.ExqCorpModeler, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	stmt, _ := db.Prepare(createUserTable)
	stmt.Exec()

	stmt, _ = db.Prepare(createCorpseTable)
	stmt.Exec()

	return &ExqCorpModelSqlite{DB: db}, nil
}

func (e *ExqCorpModelSqlite) GetCorpse(name string) (*models.Corpse, error) {
	stmt := `SELECT name, stage FROM corpses WHERE name = ?`
	row := e.QueryRow(stmt, name)

	c := &models.Corpse{}

	err := row.Scan(&c.Name, &c.Stage)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return c, nil
}
