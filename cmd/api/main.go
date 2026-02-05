package main

import (
	"errors"
	"fmt"
	"log"

	database "github.com/goblog-api/internal/infrastructure/database/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	fmt.Println("initializing project GoBlog API")

	db, err := database.NewMySQL()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db.GetConnection(), &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	defer db.Close()
}
