package db

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func newTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("mysql", "root:dbpass@tcp(172.17.0.2:3306)/test_homeserver?parseTime=true&multiStatements=true")

	if err != nil {
		t.Fatal(err)
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"test_homeserver",
		driver,
	)
	if err := m.Up(); err != nil {
		t.Fatal(err)
	}

	return db, func() {
		m.Down()
		db.Close()
	}
}
