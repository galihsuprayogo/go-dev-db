package godevdb

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3308)/go_dev_db")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
}
