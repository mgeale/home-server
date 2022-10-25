package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mgeale/homeserver/pkg/models"
	"github.com/mgeale/homeserver/pkg/models/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	balances interface {
		Insert(string, float32, float32, int, int) (int, error)
		Update(int, string, float32, float32, int, int) error
		Get(int) (*models.Balance, error)
		Delete(int) error
		Latest() ([]*models.Balance, error)
	}
	transactions interface {
		Insert(string, float32, string, string) (int, error)
		Update(int, string, float32, string, string) error
		Get(int) (*models.Transaction, error)
		Delete(int) error
		Latest() ([]*models.Transaction, error)
	}
	users interface {
		Authenticate(string, string) error
		Get(int) (*models.User, error)
	}
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/homeserver?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		balances:     &mysql.BalanceModel{DB: db},
		transactions: &mysql.TransactionModel{DB: db},
		users:        &mysql.UserModel{DB: db},
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
