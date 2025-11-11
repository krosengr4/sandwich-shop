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

func (d *Database) GetAllOrders() ([]*models.Order, error) {
	query := "SELECT order_id, customer_name, quantity_of_items, total_price, time_ordered FROM orders;"

	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %w", err)
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		var order models.Order

		err := rows.Scan(&order.ID, &order.CustomerName, &order.Quantity, &order.TotalPrice, &order.TimeOfOrder)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return orders, nil
}

func (d *Database) AddOrder(order *models.Order) error {
	query := "INSERT INTO orders (customer_name, quantity_of_items, total_price, time_ordered) VALUES (?, ?, ?, ?);"

	_, err := d.conn.Exec(query, order.CustomerName, order.Quantity, order.TotalPrice, order.TimeOfOrder)
	if err != nil {
		return fmt.Errorf("failed to add the order: %w", err)
	}

	fmt.Println("Success! The order was added to the db!")
	return nil
}

func (d *Database) EditOrder(newOrder *models.Order) error {
	query := "UPDATE orders SET customer_name = ?, quantity_of_items = ?, total_price = ? WHERE order_id = ?"

	_, err := d.conn.Exec(query, editField, newValue, orderId
	if err != nil {
		return fmt.Errorf("failed to update the order: %w", err)
	}

	newOrder.PrintData()
	fmt.Println("Success! The order was updated!")
	return nil
}

