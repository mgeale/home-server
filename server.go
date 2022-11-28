package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mgeale/homeserver/graph"
	"github.com/mgeale/homeserver/graph/generated"
	"github.com/mgeale/homeserver/internal/app"
	"github.com/mgeale/homeserver/internal/db"
	"github.com/mgeale/homeserver/internal/jsonlog"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

	database, err := openDB("root:dbpass@(172.17.0.2:3306)/test_homeserver?parseTime=true")
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer database.Close()

	app := &app.Application{
		Logger: logger,
		Models: db.NewModels(database),
	}
	resolvers := graph.NewResolver(app)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
