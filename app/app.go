package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Application struct {
	Db *sql.DB
}

func (app *Application) load() error {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	app.Db = db
	fmt.Println("Successfully connected!")

	return nil
}

func (app *Application) unload() {
	app.Db.Close()
}

func (app *Application) start() {
	http.HandleFunc("/token", app.issueToken)
	http.HandleFunc("/login", app.getLoginByToken)

	port := os.Getenv("APP_PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Failed to start server; ", err)
	}
}
