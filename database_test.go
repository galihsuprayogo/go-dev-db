package godevdb

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("mysql", "")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
}
