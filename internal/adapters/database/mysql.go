package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) (*MySQL, error) {
	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		return nil, fmt.Errorf("env DB_DSN not found")
	}

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &MySQL{conn: conn}, nil
}

func (m *MySQL) Close() error {
	err := m.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}

func (m *MySQL) GetConnection() *sql.DB {
	return m.conn
}
