package database

import (
	"database/sql"
	"fmt"
	"sandwich-shop/config"
)

type Database struct {
	conn *sql.DB
}

func GetConnection(cfg *config.DatabaseConfig) (*Database, error) {
	// Create connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	// Open sql database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %w", err)
	}

	// Ping(verify connection to db) sql database
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{conn: db}, nil
}

func (d *Database) Close() error {
	return d.conn.Close()
}

// todo: Method that gets all orders from the db

// todo: Method that adds an order to the db
