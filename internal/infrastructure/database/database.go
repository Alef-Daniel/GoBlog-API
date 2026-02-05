package database

import "database/sql"

type Database interface {
	GetConnection() *sql.DB
}
