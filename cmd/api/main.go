package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mgeale/homeserver/internal/app"
	"github.com/mgeale/homeserver/internal/db"
	"github.com/mgeale/homeserver/internal/jsonlog"
)

func main() {
	var cfg app.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production")

	pw := os.Getenv("DB_PW")
	flag.StringVar(&cfg.Db.Dsn, "db-dsn", fmt.Sprintf("web:%s@(localhost:3306)/homeserver?parseTime=true", pw), "MySQL DSN")

	flag.IntVar(&cfg.Db.MaxOpenConns, "db-max-open-conns", 25, "MySQL max open connections")
	flag.IntVar(&cfg.Db.MaxIdleConns, "db-max-idle-conns", 25, "MySQL max open idle connections")
	flag.StringVar(&cfg.Db.MaxIdleTime, "db-max-idle-time", "15m", "MySQL max connection idle time")

	flag.Float64Var(&cfg.Limiter.Rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()

	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

	database, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer database.Close()

	app := &app.Application{
		Logger: logger,
		Config: cfg,
		Models: db.NewModels(database),
	}

	if err := serve(app); err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg app.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Db.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Db.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
