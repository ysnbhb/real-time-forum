package module

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Init() (*sql.DB, error) {
	dbCommand, err := os.ReadFile("./database/codesql.sql")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db, err := sql.Open("sqlite3", "./database/real.db")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	_, err = db.Exec(string(dbCommand))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, err
}
