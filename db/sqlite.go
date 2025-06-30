package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteDb struct {
    db *sql.DB
}

func OpenSqliteDatabase(dataSource string) (Database ,error) {
	db, err := sql.Open("sqlite3", dataSource)
    var sqDb *sqliteDb = nil;

    if db != nil {
        _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tags (
            tag TEXT PRIMARY KEY,
            location TEXT NOT NULL
        )`)
        
        if err != nil {
            return sqDb, err
        }
        
        sqDb = &sqliteDb{
            db: db,
        }
    }

    return sqDb, err
}

func (sqDb sqliteDb) AddTag(tag string, location string) error {
    _, err := sqDb.db.Exec(`INSERT INTO tags VALUES (?, ?)`, tag, location)
    return err
}

func (sqDb sqliteDb) FetchTag(tag string) (string, error) {
    var location string
    row := sqDb.db.QueryRow(`SELECT location FROM tags WHERE tag = ? COLLATE NOCASE`, tag)
    err := row.Scan(&location)
    
    if err != nil {
        if err == sql.ErrNoRows {
            err = &TagNotFoundErr{tag: ""}
        }
        return "", err
    }
    
    return location, nil
}

func (sqDb sqliteDb) ListTags(prefix string) ([]Tag, error) {
    prefix += "%"
    rows, err := sqDb.db.Query(`SELECT tag, location 
        FROM tags WHERE 
        tag LIKE ? COLLATE NOCASE`, prefix)
    if err != nil {
        return nil, err
    }
    
    defer rows.Close()
    
    tags := []Tag{}
    for rows.Next() {
        var tag Tag
        rows.Scan(&tag.tag, &tag.location)
        tags = append(tags, tag)
    }
    
    return tags, nil
}

func (sqDb sqliteDb) Close() error {
    return sqDb.db.Close()
}