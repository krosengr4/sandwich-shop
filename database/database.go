package database

import (
	"database/sql"
	"fmt"
	"sandwich-shop/config"
	"sandwich-shop/models"
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
func (d *Database) GetAllOrders() ([]*models.Order, error) {
	query := "SELECT * FROM orders;"

	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %w", err)
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		var order models.Order

		err := rows.Scan(&order.ID, order.CustomerName, order.ItemsOrdered, order.TotalPrice, order.TotalPrice, order.TimeOfOrder)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

// todo: Method that adds an order to the db
