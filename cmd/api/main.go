package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mgeale/homeserver/internal/app"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &app.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	if err := serve(app); err != nil {
		// logger.PrintFatal(err, nil)
	}
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:dbpass@(172.17.0.2:3306)/homeserver?parseTime=true")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
