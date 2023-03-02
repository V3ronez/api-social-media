package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver (implicit import)
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	fmt.Println("⚡️ Database running... 🔥")
	return db, nil
}
